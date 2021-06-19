package merchant

import (
	"context"
	"errors"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateMerchant(ctx context.Context, args []string) error
	UpdateMerchantDiscount(ctx context.Context, args []string) error
}

type service struct {
	r Repository
}

func NewMerchantService(repo Repository) Service {
	return &service{
		r: repo,
	}
}

func (s *service) CreateMerchant(ctx context.Context, args []string) error {

	var (
		merchantDetails types.MerchantDetails
	)
	if len(args) != 3 {
		return errors.New(constants.CreateMerchantInvalidParamsErr)
	}

	perc, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return errors.New(constants.CreateMerchantPercErr)
	}

	if perc <= 0 || perc > 100 {
		return errors.New(constants.CreateMerchantInvalidDiscountErr)
	}

	merchantDetails.Name = args[0]
	merchantDetails.Email = args[1]
	merchantDetails.Perc = perc

	return s.r.CreateMerchant(ctx, merchantDetails)
}

func (s *service) UpdateMerchantDiscount(ctx context.Context, args []string) error {

	if len(args) != 2 {
		return errors.New(constants.UpdateMerchantInvalidParamsErr)
	}

	name := args[0]
	if len(name) == 0 {
		return errors.New(constants.UpdateMerchantNameMissingErr)
	}

	discount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return errors.New(constants.UpdateMerchantPercErr)
	}

	// If discount is 0 or less or greater that 100, its an invalid discount
	if discount <= 0 || discount > 100 {
		return errors.New(constants.UpdateMerchantInvalidDiscountErr)
	}

	return s.r.UpdateMerchant(ctx, name, map[string]interface{}{
		"perc": discount,
	})
}
