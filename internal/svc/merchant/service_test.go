package merchant

import (
	"context"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateMerchant_Success(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("CreateMerchant", mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.CreateMerchant(context.Background(), []string{"m1", "m1@gmail.com", "1.5"})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceCreateMerchant_InvalidArgs(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("CreateMerchant", mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.CreateMerchant(context.Background(), []string{})
	assert.EqualError(t, err, constants.CreateMerchantInvalidParamsErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceCreateMerchant_InvalidDiscount(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("CreateMerchant", mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.CreateMerchant(context.Background(), []string{"u1", "u1@gmail.com", "hundred"})
	assert.EqualError(t, err, constants.CreateMerchantPercErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceCreateMerchant_DiscountGreaterThan100(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("CreateMerchant", mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.CreateMerchant(context.Background(), []string{"u1", "u1@gmail.com", "101"})
	assert.EqualError(t, err, constants.CreateMerchantInvalidDiscountErr)
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestServiceUpdateMerchantDiscount_Success(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("UpdateMerchant", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.UpdateMerchantDiscount(context.Background(), []string{"m1", "1.5"})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceUpdateMerchantDiscount_InvalidArgs(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("UpdateMerchant", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.UpdateMerchantDiscount(context.Background(), []string{})
	assert.EqualError(t, err, constants.UpdateMerchantInvalidParamsErr)
	mockRepo.AssertNotCalled(t, "UpdateMerchant", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateMerchantDiscount_NameIsEmpty(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("UpdateMerchant", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.UpdateMerchantDiscount(context.Background(), []string{"", "10"})
	assert.EqualError(t, err, constants.UpdateMerchantNameMissingErr)
	mockRepo.AssertNotCalled(t, "UpdateMerchant", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateMerchantDiscount_InvalidDiscount(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("UpdateMerchant", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.UpdateMerchantDiscount(context.Background(), []string{"m1", "abc"})
	assert.EqualError(t, err, constants.UpdateMerchantPercErr)
	mockRepo.AssertNotCalled(t, "UpdateMerchant", mock.Anything, mock.Anything, mock.Anything)
}

func TestServiceUpdateMerchantDiscount_DiscountLessThan0(t *testing.T) {

	mockRepo := getMerchantRepoMock()
	mockRepo.On("UpdateMerchant", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewMerchantService(mockRepo)
	err := svc.UpdateMerchantDiscount(context.Background(), []string{"m1", "-1"})
	assert.EqualError(t, err, constants.UpdateMerchantInvalidDiscountErr)
	mockRepo.AssertNotCalled(t, "UpdateMerchant", mock.Anything, mock.Anything, mock.Anything)
}
