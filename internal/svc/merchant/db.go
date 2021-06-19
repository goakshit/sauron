package merchant

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Repository interface {
	CreateMerchant(ctx context.Context, data types.MerchantDetails) error
	UpdateMerchant(ctx context.Context, name string, update map[string]interface{}) error
}

type repository struct {
	db persistence.DBIface
}

func NewRepository(db persistence.DBIface) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateMerchant(ctx context.Context, data types.MerchantDetails) error {
	return r.db.Table("merchant").Create(&data).Error()
}

func (r *repository) UpdateMerchant(ctx context.Context, name string, update map[string]interface{}) error {
	res := r.db.Table("merchant").Where("name = ?", name).Updates(update)
	if res.Error() == nil && res.RowsAffected() != 0 {
		return nil
	} else if res.Error() == nil {
		return errors.New(constants.UpdateMerchantNotFoundErr)
	} else {
		return res.Error()
	}
}
