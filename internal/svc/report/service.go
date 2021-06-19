package report

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"gorm.io/gorm"
)

type Service interface {
	GetUsersAtCreditLimit(ctx context.Context) ([]string, error)
	GetUserDues(ctx context.Context, name string) (float64, error)
	GetTotalDues(ctx context.Context) ([]types.ReportUserDues, error)
	GetMerchantDiscount(ctx context.Context, name string) (float64, error)
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

func (s *service) GetTotalDues(ctx context.Context) ([]types.ReportUserDues, error) {

	users := []types.ReportUserDues{}
	err := s.db.Table("user").Where("due_amount > 0").Find(&users).Error()
	if err != nil {
		return users, errors.New(constants.ReportUserDuesGetUsersErr)
	}
	return users, nil
}

func (s *service) GetUserDues(ctx context.Context, name string) (float64, error) {

	user := types.ReportUserDues{}
	err := s.db.Table("user").Where("name = ?", name).First(&user).Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.New(constants.ReportUserDuesUserNotFoundErr)
		}
		return 0, errors.New(constants.ReportUserDuesGetUserDueErr)
	}
	return user.DueAmount, nil
}

func (s *service) GetMerchantDiscount(ctx context.Context, name string) (float64, error) {

	txns := []types.ReportMerchantTxn{}
	err := s.db.Table("transaction").Where("merchant_name = ?", name).Find(&txns).Error()
	if err != nil {
		return 0, errors.New(constants.ReportDiscountGetTxnErr)
	}
	if len(txns) == 0 {
		return 0, nil
	}
	var result float64
	for _, txn := range txns {
		discount := (txn.MerchantPerc * txn.Amount) / 100
		result += discount
	}
	return result, nil
}
