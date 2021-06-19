package transaction

import (
	"context"
	"errors"
	"testing"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction_Success(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   0,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	}, nil)
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(nil)
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u1",
		MerchantName: "m1",
		Amount:       500,
	})
	assert.Nil(t, err)
	txnRepoMock.AssertExpectations(t)
}

func TestCreateTransaction_UserNotFound(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{}, errors.New(constants.CreateTxnUserDoesNotExistErr))
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	}, nil)
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(nil)
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u-not-exist",
		MerchantName: "m1",
		Amount:       500,
	})
	assert.EqualError(t, err, constants.CreateTxnUserDoesNotExistErr)
	txnRepoMock.AssertNotCalled(t, "GetMerchantDetails", mock.Anything, mock.Anything)
	txnRepoMock.AssertNotCalled(t, "CreateTxn", mock.Anything, mock.Anything)
	txnRepoMock.AssertNotCalled(t, "UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreateTransaction_MerchantNotFound(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   0,
		CreditLimit: 1000,
	}, nil)
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{}, errors.New(constants.CreateTxnMerchantNotFoundErr))
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(nil)
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u1",
		MerchantName: "m-not-exist",
		Amount:       500,
	})
	assert.EqualError(t, err, constants.CreateTxnMerchantNotFoundErr)
	txnRepoMock.AssertNotCalled(t, "CreateTxn", mock.Anything, mock.Anything)
	txnRepoMock.AssertNotCalled(t, "UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreateTransaction_CreditLimitExceeded(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   200,
		CreditLimit: 500,
	}, nil)
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	}, nil)
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(nil)
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u1",
		MerchantName: "m1",
		Amount:       301,
	})
	assert.EqualError(t, err, constants.CreateTxnUserCreditLimitExceededErr)
	txnRepoMock.AssertNotCalled(t, "CreateTxn", mock.Anything, mock.Anything)
	txnRepoMock.AssertNotCalled(t, "UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreateTransaction_CreateTxnFailed(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   200,
		CreditLimit: 500,
	}, nil)
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	}, nil)
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(errors.New("Internal Server Error"))
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u-not-exist",
		MerchantName: "m1",
		Amount:       300,
	})
	assert.EqualError(t, err, "Internal Server Error")
	txnRepoMock.AssertNotCalled(t, "UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreateTransaction_UpdateUserDueAmountFailed(t *testing.T) {
	txnRepoMock := getTransactionMock()
	txnRepoMock.On("GetUserDetails", mock.Anything, mock.Anything).Return(types.UserDetails{
		Name:        "u1",
		Email:       "u1@gmail.com",
		DueAmount:   200,
		CreditLimit: 500,
	}, nil)
	txnRepoMock.On("GetMerchantDetails", mock.Anything, mock.Anything).Return(types.MerchantDetails{
		Name:  "m1",
		Email: "m1@gmail.com",
		Perc:  1.5,
	}, nil)
	txnRepoMock.On("CreateTxn", mock.Anything, mock.Anything).Return(nil)
	txnRepoMock.On("UpdateUserDueAmount", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Internal Server Error"))
	svc := NewTxnService(txnRepoMock)
	err := svc.CreateTransaction(context.Background(), types.TxnDetails{
		ID:           0,
		UserName:     "u-not-exist",
		MerchantName: "m1",
		Amount:       300,
	})
	assert.EqualError(t, err, "Internal Server Error")
	txnRepoMock.AssertExpectations(t)
}
