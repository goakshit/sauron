package transaction

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"gorm.io/gorm"
)

type Service interface {
	CreateTransaction(ctx context.Context, data types.TxnDetails) error
}

type service struct {
	db persistence.DBIface
}

func NewTxnService(db persistence.DBIface) Service {
	return &service{
		db: db,
	}
}

// CreateTransaction - Checks if credit limit is not execeeded & if not, then creates transaction and updates the due
// amount in user table also
func (s *service) CreateTransaction(ctx context.Context, data types.TxnDetails) error {

	var (
		userDetails   types.UserDetails
		merchantCount int64
	)

	if data.Amount <= 0 {
		return errors.New(constants.CreateTxnInvalidAmountErr)
	}

	err := s.db.Table("user").Where("name = ?", data.UserName).First(&userDetails).Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(constants.CreateTxnUserDoesNotExistErr)
		}
		return errors.New(constants.CreateTxnGetUserCreditLimitErr)
	}

	err = s.db.Table("merchant").Where("name = ?", data.MerchantName).Count(&merchantCount).Error()
	if err != nil {
		return errors.New(constants.CreateTxnGetMerchantErr)
	}

	if merchantCount == 0 {
		return errors.New(constants.CreateTxnMerchantNotFoundErr)
	}

	// If txn amount is greater than limit - dueAmount, throw error
	if data.Amount > (userDetails.CreditLimit - userDetails.DueAmount) {
		return errors.New(constants.CreateTxnUserCreditLimitExceededErr)
	}

	if err = s.db.Table("transaction").Create(&data).Error(); err == nil {
		// Update due amount in user table
		return s.db.Table("user").Where("name = ?", data.UserName).
			UpdateColumn("due_amount", userDetails.DueAmount+data.Amount).Error()
	}
	return err
}
