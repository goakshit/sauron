package transaction

import (
	"context"

	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/mock"
)

type transactionMock struct {
	mock.Mock
}

// GetTransactionMock - Returns service mock
// Mocks repository interface
func getTransactionMock() *transactionMock {
	return &transactionMock{}
}

func (m *transactionMock) GetUserDetails(ctx context.Context, name string) (types.UserDetails, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(types.UserDetails), args.Error(1)
}

func (m *transactionMock) GetMerchantDetails(ctx context.Context, name string) (types.MerchantDetails, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(types.MerchantDetails), args.Error(1)
}

func (m *transactionMock) CreateTxn(ctx context.Context, data types.TxnDetails) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *transactionMock) UpdateUserDueAmount(ctx context.Context, name string, dueAmount float64) error {
	args := m.Called(ctx, name, dueAmount)
	return args.Error(0)
}
