package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateUser(ctx context.Context, args []string) error
	UpdateUserCreditLimit(ctx context.Context, args []string) error
}

type service struct {
	r Repository
}

func NewUserService(repo Repository) Service {
	return &service{
		r: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, args []string) error {

	var userDetails types.UserDetails

	if len(args) != 3 {
		return errors.New(constants.CreateUserInvalidParamsErr)
	}

	creditLimit, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return errors.New(constants.CreateUserInvalidCreditLimitErr)
	}

	if creditLimit <= 0 {
		return errors.New(constants.CreateUserInvalidCreditLimitErr)
	}

	userDetails.Name = args[0]
	userDetails.Email = args[1]
	userDetails.CreditLimit = creditLimit
	return s.r.CreateUser(ctx, userDetails)
}

func (s *service) UpdateUserCreditLimit(ctx context.Context, args []string) error {

	if len(args) != 2 {
		return errors.New(constants.UpdateUserInvalidParamsErr)
	}

	creditLimit, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return errors.New(constants.UpdateUserInvalidCreditLimitErr)
	}

	if creditLimit < 0 {
		return errors.New(constants.UpdateUserInvalidCreditLimitErr)
	}

	name := args[0]
	// If name is empty, return error.
	if len(name) == 0 {
		return errors.New(constants.UpdateUserNameMissingErr)
	}
	return s.r.UpdateUserCreditLimit(ctx, name, creditLimit)
}
