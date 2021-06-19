package transaction

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateTransaction(ctx context.Context, data types.TxnDetails) error
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
func (s *service) CreateTransaction(ctx context.Context, data types.TxnDetails) error {

	var (
		userDetails     types.UserDetails
		merchantDetails types.MerchantDetails
		err             error
	)

	if data.Amount <= 0 {
		return errors.New(constants.CreateTxnInvalidAmountErr)
	}

	userDetails, err = s.r.GetUserDetails(ctx, data.UserName)
	if err != nil {
		return err
	}

	merchantDetails, err = s.r.GetMerchantDetails(ctx, data.MerchantName)
	if err != nil {
		return err
	}

	// If txn amount is greater than limit - dueAmount, throw error
	if data.Amount > (userDetails.CreditLimit - userDetails.DueAmount) {
		return errors.New(constants.CreateTxnUserCreditLimitExceededErr)
	}

	// Set current merchant discount percent in txn
	data.MerchantPerc = merchantDetails.Perc

	if err = s.r.CreateTxn(ctx, data); err == nil {
		// Update due amount in user table
		return s.r.UpdateUserDueAmount(ctx, data.UserName, userDetails.DueAmount+data.Amount)
	}
	return err
}
