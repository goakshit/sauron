package merchant

import (
	"context"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateMerchant(ctx context.Context, data types.MerchantDetails) error
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
