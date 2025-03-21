package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DingBao-sys/simple_bank/utils"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	owner := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    owner.Username,
		Balance:  utils.GenerateRandomMoney(),
		Currency: utils.GenerateRandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	require.NotEmpty(t, account)

	testAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testAccount)
	require.Equal(t, account.ID, testAccount.ID)
	require.Equal(t, account.Balance, testAccount.Balance)
	require.Equal(t, account.Owner, testAccount.Owner)
	require.Equal(t, account.Currency, testAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, testAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.GenerateRandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, readErr := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, readErr)
	require.EqualError(t, readErr, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	numOfAccounts := 10

	for i := 0; i < numOfAccounts; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  "user",
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 0)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
