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

func TestCreateUserSuccess(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewUserService(mockRepo)
	err := svc.CreateUser(context.Background(), types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		CreditLimit: 500,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUserInternalServerError(t *testing.T) {

	err := errors.New("Something went wrong")

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(err)

	svc := NewUserService(mockRepo)
	createUserErr := svc.CreateUser(context.Background(), types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		CreditLimit: 500,
	})
	assert.Errorf(t, err, createUserErr.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimitSuccess(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)
	mockRepo.On("RowsAffected").Return(int64(1))

	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", 600)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_UserDoesNotExist(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)
	mockRepo.On("RowsAffected").Return(int64(0))

	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", 600)
	assert.NotNil(t, err)
	assert.EqualError(t, err, constants.UpdateUserNotFoundErr)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_InternalServerErr(t *testing.T) {

	mockErr := errors.New("Something went wrong. Internal server error.")

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("UpdateColumn", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(mockErr)

	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", 600)
	assert.NotNil(t, err)
	assert.EqualError(t, err, mockErr.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserCreditLimit_UserNameIsEmpty(t *testing.T) {

	mockRepo := persistence.GetGormMock()

	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "", 600)
	assert.NotNil(t, err)
	assert.EqualError(t, err, constants.UpdateUserNameMissingErr)
}

func TestUpdateUserCreditLimit_InvalidCreditLimit(t *testing.T) {

	mockRepo := persistence.GetGormMock()

	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), "u1", -150)
	assert.NotNil(t, err)
	assert.EqualError(t, err, constants.UpdateUserInvalidCreditLimitErr)
}
