package transaction

import (
	"context"
	"errors"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateTransaction(ctx context.Context, args []string) error
}

type service struct {
	r Repository
}

func NewTxnService(repo Repository) Service {
	return &service{
		r: repo,
	}
}

// CreateTransaction - Checks if credit limit is not execeeded & if not, then creates transaction and updates the due
// amount in user table also
func (s *service) CreateTransaction(ctx context.Context, args []string) error {

	var (
		txnDetails      types.TxnDetails
		userDetails     types.UserDetails
		merchantDetails types.MerchantDetails
		err             error
	)

	if len(args) != 3 {
		return errors.New(constants.CreateTxnInvalidParamsErr)
	}

	txnAmount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return errors.New(constants.CreateTxnInvalidAmountErr)
	}

	if txnAmount <= 0 {
		return errors.New(constants.CreateTxnInvalidAmountErr)
	}

	txnDetails.UserName = args[0]
	txnDetails.MerchantName = args[1]
	txnDetails.Amount = txnAmount

	userDetails, err = s.r.GetUserDetails(ctx, txnDetails.UserName)
	if err != nil {
		return err
	}

	merchantDetails, err = s.r.GetMerchantDetails(ctx, txnDetails.MerchantName)
	if err != nil {
		return err
	}

	// If txn amount is greater than limit - dueAmount, throw error
	if txnDetails.Amount > (userDetails.CreditLimit - userDetails.DueAmount) {
		return errors.New(constants.CreateTxnUserCreditLimitExceededErr)
	}

	// Set current merchant discount percent in txn
	txnDetails.MerchantPerc = merchantDetails.Perc

	if err = s.r.CreateTxn(ctx, txnDetails); err == nil {
		// Update due amount in user table
		return s.r.UpdateUserDueAmount(ctx, txnDetails.UserName, userDetails.DueAmount+txnDetails.Amount)
	}
	return err
}
