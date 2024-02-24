package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func CreateRandomAccount(t *testing.T) (Accounts, CreateAccountParams, error) {
	user, _, _ := CreateRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	return account, arg, err
}

func TestCreateAccount(t *testing.T) {
	account, arg, err := CreateRandomAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func TestGetAccountById(t *testing.T) {
	account, _, err := CreateRandomAccount(t)
	require.NoError(t, err)

	result, err := testQueries.GetAccountByID(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.Owner, account.Owner)
	require.Equal(t, result.Balance, account.Balance)
	require.Equal(t, result.Currency, account.Currency)
	require.WithinDuration(t, result.CreatedAt, account.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account, _, err := CreateRandomAccount(t)
	require.NoError(t, err)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomBalance(),
	}

	result, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, result.Balance, account.Balance)

	require.Equal(t, account.ID, result.ID)
	require.Equal(t, account.Currency, result.Currency)
	require.Equal(t, account.Owner, result.Owner)
	require.WithinDuration(t, result.CreatedAt, account.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account, _, _ := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	result, err := testQueries.GetAccountByID(context.Background(), account.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, result)
}

func TestGetAllAccounts(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		CreateRandomAccount(t)
	}

	arg := GetAllAccountsParams{
		Limit:  int32(n / 2),
		Offset: int32(n / 2),
	}

	accounts, err := testQueries.GetAllAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.Len(t, accounts, n/2)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
