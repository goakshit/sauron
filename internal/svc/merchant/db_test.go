package merchant

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

func TestCreateMerchant_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.CreateMerchant(context.Background(), types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateMerchant_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.CreateMerchant(context.Background(), types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateMerchant_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Updates", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("RowsAffected").Return(int64(1))
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateMerchant(context.Background(), "m1", map[string]interface{}{
		"perc": 1.25,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMerchant_MerchantNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Updates", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("RowsAffected").Return(int64(0))
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateMerchant(context.Background(), "m-not-exist", map[string]interface{}{
		"perc": 1.25,
	})
	assert.EqualError(t, err, constants.UpdateMerchantNotFoundErr)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMerchant_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Updates", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.UpdateMerchant(context.Background(), "m1", map[string]interface{}{
		"perc": 1.25,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}
