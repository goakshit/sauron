package user

import (
	"context"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
)

type Service interface {
	CreateUser(ctx context.Context, data types.UserDetails) error
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
