package transaction

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserDetails(ctx context.Context, name string) (types.UserDetails, error)
	GetMerchantDetails(ctx context.Context, name string) (types.MerchantDetails, error)
	CreateTxn(ctx context.Context, data types.TxnDetails) error
	UpdateUserDueAmount(ctx context.Context, name string, dueAmount float64) error
}

type repository struct {
	db persistence.DBIface
}

func NewRepository(db persistence.DBIface) Repository {
	return &repository{
		db: db,
	}
}

// GetUserDetails - Fetches user details by name
func (r *repository) GetUserDetails(ctx context.Context, name string) (types.UserDetails, error) {
	var (
		userDetails types.UserDetails
	)
	err := r.db.Table("user").Where("name = ?", name).First(&userDetails).Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userDetails, errors.New(constants.CreateTxnUserDoesNotExistErr)
		}
		return userDetails, errors.New(constants.CreateTxnGetUserCreditLimitErr)
	}
	return userDetails, nil
}

// GetMerchantDetails - Fetches merchant details by name
func (r *repository) GetMerchantDetails(ctx context.Context, name string) (types.MerchantDetails, error) {
	var (
		merchantDetails types.MerchantDetails
	)
	err := r.db.Table("merchant").Where("name = ?", name).First(&merchantDetails).Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return merchantDetails, errors.New(constants.CreateTxnMerchantNotFoundErr)
		}
		return merchantDetails, errors.New(constants.CreateTxnGetMerchantErr)
	}
	return merchantDetails, nil
}

// CreateTxn - Creates txn record in db
func (r *repository) CreateTxn(ctx context.Context, data types.TxnDetails) error {
	return r.db.Table("transaction").Create(&data).Error()
}

// UpdateUserDueAmount - Updates user due amount
func (r *repository) UpdateUserDueAmount(ctx context.Context, name string, dueAmount float64) error {
	return r.db.Table("user").Where("name = ?", name).UpdateColumn("due_amount", dueAmount).Error()
}
