package user

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.CreateUser(context.Background(), types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		CreditLimit: 500,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.CreateUser(context.Background(), types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		CreditLimit: 500,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("RowsAffected").Return(int64(1))
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", 500)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_RecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("RowsAffected").Return(int64(0))
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u-not-found", 500)
	assert.EqualError(t, err, constants.UpdateUserNotFoundErr)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", 500)
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}
