package payback

import (
	"context"

	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/mock"
)

type paybackMock struct {
	mock.Mock
}

// getPaybackMock - Returns service mock
// Mocks repository interface
func getPaybackMock() *paybackMock {
	return &paybackMock{}
}

func (m *paybackMock) CreatePayback(ctx context.Context, data types.PaybackDetails) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *paybackMock) UpdateUser(ctx context.Context, name string, update map[string]interface{}) error {
	args := m.Called(ctx, name, update)
	return args.Error(0)
}
