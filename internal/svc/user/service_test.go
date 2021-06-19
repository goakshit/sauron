package user

import (
	"context"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateUser_Success(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.CreateUser(context.Background(), []string{"u1", "u1@gmail.com", "500"})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceCreateUser_InvalidArgs(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.CreateUser(context.Background(), []string{})
	assert.EqualError(t, err, constants.CreateUserInvalidParamsErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceCreateUser_InvalidCreditLimit(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.CreateUser(context.Background(), []string{"u1", "u1@gmail.com", "hundred"})
	assert.EqualError(t, err, constants.CreateUserInvalidCreditLimitErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceCreateUser_CreditLimitLessThan0(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.CreateUser(context.Background(), []string{"u1", "u1@gmail.com", "-1"})
	assert.EqualError(t, err, constants.CreateUserInvalidCreditLimitErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceUpdateUserCreditLimit_Success(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), []string{"u1", "500"})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceUpdateUserCreditLimit_InvalidArgs(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), []string{})
	assert.EqualError(t, err, constants.UpdateUserInvalidParamsErr)
	mockRepo.AssertNotCalled(t, "UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateUserCreditLimit_InvalidCreditLimit(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), []string{"u1", "abc"})
	assert.EqualError(t, err, constants.UpdateUserInvalidCreditLimitErr)
	mockRepo.AssertNotCalled(t, "UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateUserCreditLimit_CreditLimitLessThan0(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), []string{"u1", "-100"})
	assert.EqualError(t, err, constants.UpdateUserInvalidCreditLimitErr)
	mockRepo.AssertNotCalled(t, "UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateUserCreditLimit_EmptyName(t *testing.T) {

	mockRepo := getUserRepoMock()
	mockRepo.On("UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewUserService(mockRepo)
	err := svc.UpdateUserCreditLimit(context.Background(), []string{"", "500"})
	assert.EqualError(t, err, constants.UpdateUserNameMissingErr)
	mockRepo.AssertNotCalled(t, "UpdateUserCreditLimit", mock.Anything, mock.Anything, mock.Anything)
}
