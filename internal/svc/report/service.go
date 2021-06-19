package report

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
)

type Service interface {
	GetUsersAtCreditLimit(ctx context.Context) ([]string, error)
}

type service struct {
	db persistence.DBIface
}

func NewReportService(db persistence.DBIface) Service {
	return &service{
		db: db,
	}
}

func (s *service) GetUsersAtCreditLimit(ctx context.Context) ([]string, error) {

	users := []string{}
	err := s.db.Table("user").Select("name").Where("due_amount = credit_limit").Find(&users).Error()
	if err != nil {
		return users, errors.New(constants.ReportUACLGetUsersErr)
	}
	return users, nil
}
