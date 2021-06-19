package user

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Repository interface {
	CreateUser(ctx context.Context, data types.UserDetails) error
	UpdateUserCreditLimit(ctx context.Context, name string, creditLimit float64) error
}

type repository struct {
	db persistence.DBIface
}

func NewRepository(db persistence.DBIface) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, data types.UserDetails) error {
	return r.db.Table("user").Create(&data).Error()
}

func (r *repository) UpdateUserCreditLimit(ctx context.Context, name string, creditLimit float64) error {

	res := r.db.Table("user").Where("name = ?", name).UpdateColumn("credit_limit", creditLimit)
	if res.Error() == nil && res.RowsAffected() != 0 {
		return nil
	} else if res.Error() == nil {
		return errors.New(constants.UpdateUserNotFoundErr)
	} else {
		return res.Error()
	}
}
