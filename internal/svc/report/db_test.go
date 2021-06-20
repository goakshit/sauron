package report

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetUsersAtCreditLimit_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Select", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetUsersAtCreditLimit(context.Background())
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUsersAtCreditLimit_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Select", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetUsersAtCreditLimit(context.Background())
	assert.EqualError(t, err, constants.ReportUACLGetUsersErr)
	mockRepo.AssertExpectations(t)
}

func TestGetTotalDues_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetTotalDues(context.Background())
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTotalDues_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetTotalDues(context.Background())
	assert.EqualError(t, err, constants.ReportUserDuesGetUsersErr)
	mockRepo.AssertExpectations(t)
}

func TestGetUserDues_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDues(context.Background(), "u1")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserDues_UserRecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(gorm.ErrRecordNotFound)

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDues(context.Background(), "u-not-exist")
	assert.EqualError(t, err, constants.ReportUserDuesUserNotFoundErr)
	mockRepo.AssertExpectations(t)
}

func TestGetUserDues_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDues(context.Background(), "u1")
	assert.EqualError(t, err, constants.ReportUserDuesGetUserDueErr)
	mockRepo.AssertExpectations(t)
}

func TestGetMerchantDiscount_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetMerchantDiscount(context.Background(), "m1")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetMerchantDiscount_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Find", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal server error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetMerchantDiscount(context.Background(), "m1")
	assert.EqualError(t, err, constants.ReportDiscountGetTxnErr)
	mockRepo.AssertExpectations(t)
}
