package merchant

import (
	"context"

	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/mock"
)

type merchantRepoMock struct {
	mock.Mock
}

// getmerchantRepoMock - Returns repo mock
// Mocks repository interface
func getMerchantRepoMock() *merchantRepoMock {
	return &merchantRepoMock{}
}

func (m *merchantRepoMock) CreateMerchant(ctx context.Context, data types.MerchantDetails) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *merchantRepoMock) UpdateMerchant(ctx context.Context, name string, update map[string]interface{}) error {
	args := m.Called(ctx, name, update)
	return args.Error(0)
}
