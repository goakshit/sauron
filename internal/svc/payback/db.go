package payback

import (
	"context"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Repository interface {
	CreatePayback(ctx context.Context, data types.PaybackDetails) error
	UpdateUser(ctx context.Context, name string, update map[string]interface{}) error
}

type repository struct {
	db persistence.DBIface
}

func NewRepository(db persistence.DBIface) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreatePayback(ctx context.Context, data types.PaybackDetails) error {
	return r.db.Table("payback").Create(&data).Error()
}

func (r *repository) UpdateUser(ctx context.Context, name string, update map[string]interface{}) error {
	return r.db.Table("user").Where("name = ?", name).Updates(update).Error()
}
