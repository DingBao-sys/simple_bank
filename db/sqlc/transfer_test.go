package db

import (
	"context"
	"testing"
	"time"

	"github.com/DingBao-sys/simple_bank/utils"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer {
	arg := CreateTransferParams{
		ToAccountID:   toAccount.ID,
		FromAccountID: fromAccount.ID,
		Amount:        utils.GenerateRandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	transfer := createRandomTransfer(t, fromAccount, toAccount)
	testTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testTransfer)
	require.Equal(t, transfer.ID, testTransfer.ID)
	require.Equal(t, transfer.FromAccountID, testTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, testTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, testTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, testTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	}
}
