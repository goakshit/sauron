package user

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateUser(ctx context.Context, data types.UserDetails) error
	UpdateUserCreditLimit(ctx context.Context, name string, creditLimit float64) error
}

type service struct {
	db persistence.DBIface
}

func NewUserService(db persistence.DBIface) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateUser(ctx context.Context, data types.UserDetails) error {
	return s.db.Table("user").Create(&data).Error()
}

func (s *service) UpdateUserCreditLimit(ctx context.Context, name string, creditLimit float64) error {

	// If name is empty, stop right here.
	if len(name) == 0 {
		return errors.New(constants.UpdateUserNameMissingErr)
	}

	// check for invalid credit limit
	if creditLimit < 0 {
		return errors.New(constants.UpdateUserInvalidCreditLimitErr)
	}
	res := s.db.Table("user").Where("name = ?", name).UpdateColumn("credit_limit", creditLimit)
	if res.Error() == nil && res.RowsAffected() != 0 {
		return nil
	} else if res.Error() == nil {
		return errors.New(constants.UpdateUserNotFoundErr)
	} else {
		return res.Error()
	}
}
