package payback

import (
	"context"
	"errors"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"gorm.io/gorm"
)

type Service interface {
	CreatePayback(ctx context.Context, data types.PaybackDetails) error
}

type service struct {
	db persistence.DBIface
}

func NewPaybackService(db persistence.DBIface) Service {
	return &service{
		db: db,
	}
}

// CreatePayback - Checks if user has due amount & if yes, then creates payback and updates the due
// amount - payback amount in user table also
func (s *service) CreatePayback(ctx context.Context, data types.PaybackDetails) error {

	var (
		userDetails types.UserDetails
	)

	if data.Amount <= 0 {
		return errors.New(constants.CreatePaybackInvalidAmountErr)
	}

	err := s.db.Table("user").Where("name = ?", data.UserName).First(&userDetails).Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(constants.CreatePaybackUserNotFoundErr)
		}
		return errors.New(constants.CreatePaybackGetDueAmountErr)
	}

	if userDetails.DueAmount == 0 {
		return errors.New(constants.CreatepaybackNoDueAmountErr)
	}

	// You can packback only the amount that is due.
	if data.Amount > userDetails.DueAmount {
		data.Amount = userDetails.DueAmount
	}

	if err = s.db.Table("payback").Create(&data).Error(); err == nil {
		// Update updated due amount in user table
		return s.db.Table("user").Where("name = ?", data.UserName).
			UpdateColumn("due_amount", userDetails.DueAmount-data.Amount).Error()
	}
	return err
}
