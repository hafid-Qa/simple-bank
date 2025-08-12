package db

import (
	"context"
	"database/sql"
	"go_api/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomAmount(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)
	return account
}
func TestCreateAccount(t *testing.T) {
	resetDB()
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	resetDB()

	account := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	resetDB()
	account := createRandomAccount(t)
	data := UpdateAccountParams{
		Balance: utils.RandomAmount(),
		ID:      account.ID,
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, data.ID, updatedAccount.ID)
	require.Equal(t, data.Balance, updatedAccount.Balance)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	resetDB()
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)

}

func TestListAccounts(t *testing.T) {
	resetDB()
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accs, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accs, 5)

	for _, acc := range accs {
		require.NotEmpty(t, acc)
	}

}
