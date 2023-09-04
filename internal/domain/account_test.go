package domain

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	testQueries = testQueries

	arg := createRandomAccountParams()

	acount, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acount)

	require.Equal(t, arg.Owner, acount.Owner)
	require.Equal(t, arg.Balance, acount.Balance)
	require.Equal(t, arg.Currency, acount.Currency)
	require.NotZero(t, acount.ID)
}

func createRandomAccountParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    createRandomString(),
		Balance:  createRandomNumber(),
		Currency: createRandomString(),
	}
}

func createRandomAccount(t *testing.T) Account {
	arg := createRandomAccountParams()
	account, err := testQueries.CreateAccount(context.Background(), arg)

	if err != nil {
		t.Fatal("cannot create random account")
	}

	return account
}
