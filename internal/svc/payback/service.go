package payback

import (
	"context"
	"errors"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/svc/transaction"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreatePayback(ctx context.Context, args []string) error
}

type service struct {
	pr Repository
	tr transaction.Repository
}

func NewPaybackService(pr Repository, tr transaction.Repository) Service {
	return &service{
		pr: pr,
		tr: tr,
	}
}

// CreatePayback - Checks if user has due amount & if yes, then creates payback and updates the due
// amount - payback amount in user table also
func (s *service) CreatePayback(ctx context.Context, args []string) error {

	var (
		paybackDetails types.PaybackDetails
		userDetails    types.UserDetails
	)

	if len(args) != 2 {
		return errors.New(constants.CreatePaybackInvalidParamsErr)
	}

	name := args[0]
	if len(name) == 0 {
		return errors.New(constants.CreatePaybackUserNotFoundErr)
	}

	paybackAmount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return errors.New(constants.CreatePaybackInvalidAmountErr)
	}

	if paybackAmount <= 0 {
		return errors.New(constants.CreatePaybackInvalidAmountErr)
	}

	paybackDetails.UserName = name
	paybackDetails.Amount = paybackAmount

	userDetails, err = s.tr.GetUserDetails(ctx, name)
	if err != nil {
		return err
	}

	if userDetails.DueAmount == 0 {
		return errors.New(constants.CreatepaybackNoDueAmountErr)
	}

	// You can packback only the amount that is due.
	if paybackAmount > userDetails.DueAmount {
		paybackAmount = userDetails.DueAmount
	}

	if err = s.pr.CreatePayback(ctx, paybackDetails); err == nil {
		// Update updated due amount in user table
		return s.pr.UpdateUser(ctx, name, map[string]interface{}{
			"due_amount": userDetails.DueAmount - paybackAmount,
		})
	}
	return err
}
