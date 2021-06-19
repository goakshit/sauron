package transaction

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetUserDetails_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDetails(context.Background(), "u1")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserDetails_RecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(gorm.ErrRecordNotFound)

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDetails(context.Background(), "u1")
	assert.Errorf(t, gorm.ErrRecordNotFound, err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetUserDetails_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetUserDetails(context.Background(), "u1")
	assert.Errorf(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestGetMerchantDetails_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	_, err := svc.GetMerchantDetails(context.Background(), "m1")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetMerchantDetails_RecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(gorm.ErrRecordNotFound)

	svc := NewRepository(mockRepo)
	_, err := svc.GetMerchantDetails(context.Background(), "m1")
	assert.Errorf(t, gorm.ErrRecordNotFound, err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetMerchantDetails_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	_, err := svc.GetMerchantDetails(context.Background(), "m1")
	assert.Errorf(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestCreateTxn_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.CreateTxn(context.Background(), types.TxnDetails{
		UserName:     "u1",
		MerchantName: "m1",
		MerchantPerc: 1.5,
		Amount:       100,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateTxn_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.CreateTxn(context.Background(), types.TxnDetails{
		UserName:     "u1",
		MerchantName: "m1",
		MerchantPerc: 1.5,
		Amount:       100,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserDueAmount_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateUserDueAmount(context.Background(), "u1", 150)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserDueAmount_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.UpdateUserDueAmount(context.Background(), "u1", 150)
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}
