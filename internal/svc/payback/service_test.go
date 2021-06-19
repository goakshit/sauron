package payback

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	txn "github.com/goakshit/sauron/internal/svc/transaction"
	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreatePayback_Success(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "500"})
	assert.Nil(t, err)
	txnRepoMock.AssertExpectations(t)
}

func TestCreatePayback_InvalidArgs(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{})
	assert.EqualError(t, err, constants.CreatePaybackInvalidParamsErr)
	txnRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
	paybackRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
}

func TestCreatePayback_EmptyName(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"", "20"})
	assert.EqualError(t, err, constants.CreatePaybackUserNotFoundErr)
	txnRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
	paybackRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
}

func TestCreatePayback_InvalidAmount(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "hundred"})
	assert.EqualError(t, err, constants.CreatePaybackInvalidAmountErr)
	txnRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
	paybackRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
}

func TestCreatePayback_AmountLessThan0(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   300,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "-19"})
	assert.EqualError(t, err, constants.CreatePaybackInvalidAmountErr)
	txnRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
	paybackRepoMock.AssertNotCalled(t, "GetUserDetails", mock.Anything, mock.Anything)
}

func TestCreatePayback_NoDueAmount(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	paybackRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   0,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   0,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "200"})
	assert.EqualError(t, err, constants.CreatepaybackNoDueAmountErr)
	txnRepoMock.AssertExpectations(t)
}

func TestCreatePayback_UserNotFound(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{}, errors.New(constants.CreateTxnUserDoesNotExistErr))
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "200"})
	assert.EqualError(t, err, constants.CreateTxnUserDoesNotExistErr)
	txnRepoMock.AssertExpectations(t)
}

func TestCreatePayback_CreatePaybackFailed(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   100,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(errors.New("Internal Server Error"))
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "200"})
	assert.EqualError(t, err, "Internal Server Error")
	txnRepoMock.AssertExpectations(t)
}

func TestCreatePayback_UpdateUserFailed(t *testing.T) {
	paybackRepoMock := getPaybackMock()
	txnRepoMock := txn.GetTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   100,
		CreditLimit: 1000,
	}, nil)
	paybackRepoMock.On("CreatePayback", mock.Anything, mock.Anything).Return(nil)
	paybackRepoMock.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Internal Server Error"))
	svc := NewPaybackService(paybackRepoMock, txnRepoMock)
	err := svc.CreatePayback(context.Background(), []string{"u1", "200"})
	assert.EqualError(t, err, "Internal Server Error")
	txnRepoMock.AssertExpectations(t)
	paybackRepoMock.AssertExpectations(t)
}
