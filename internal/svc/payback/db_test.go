package payback

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePayback_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.CreatePayback(context.Background(), types.PaybackDetails{
		UserName: "u1",
		Amount:   100,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreatePayback_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.CreatePayback(context.Background(), types.PaybackDetails{
		UserName: "u1",
		Amount:   100,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Updates", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewRepository(mockRepo)
	err := svc.UpdateUser(context.Background(), "u1", map[string]interface{}{
		"due_amount": 100,
	})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_InternalServerError(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Updates", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(errors.New("Internal Server Error"))

	svc := NewRepository(mockRepo)
	err := svc.UpdateUser(context.Background(), "u1", map[string]interface{}{
		"due_amount": 100,
	})
	assert.EqualError(t, err, "Internal Server Error")
	mockRepo.AssertExpectations(t)
}
