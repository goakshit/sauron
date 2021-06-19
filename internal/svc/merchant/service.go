package merchant

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateMerchant(ctx context.Context, data types.MerchantDetails) error
	UpdateMerchantDiscount(ctx context.Context, name string, discount float64) error
}

type service struct {
	db persistence.DBIface
}

func NewMerchantService(db persistence.DBIface) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateMerchant(ctx context.Context, data types.MerchantDetails) error {
	return s.db.Table("merchant").Create(&data).Error()
}

func (s *service) UpdateMerchantDiscount(ctx context.Context, name string, discount float64) error {

	// If name is empty, stop right here.
	if len(name) == 0 {
		return errors.New(constants.UpdateMerchantNameMissingErr)
	}

	// If discount is 0 or less, its an invalid discount
	if discount <= 0 {
		return errors.New(constants.UpdateMerchantInvalidDiscountErr)
	}
	res := s.db.Table("merchant").Where("name = ?", name).UpdateColumn("perc", discount)
	if res.Error() == nil && res.RowsAffected() != 0 {
		return nil
	} else if res.Error() == nil {
		return errors.New(constants.UpdateMerchantNotFoundErr)
	} else {
		return res.Error()
	}
}
