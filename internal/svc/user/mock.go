package user

import (
	"context"

	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

// getUserRepoMock - Returns repo mock
// Mocks repository interface
func getUserRepoMock() *userRepoMock {
	return &userRepoMock{}
}

func (m *userRepoMock) CreateUser(ctx context.Context, data types.UserDetails) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *userRepoMock) UpdateUserCreditLimit(ctx context.Context, name string, creditLimit float64) error {
	args := m.Called(ctx, name, creditLimit)
	return args.Error(0)
}
