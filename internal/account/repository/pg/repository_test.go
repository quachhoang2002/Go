package pg

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/quachhoang2002/Go/internal/domain"
	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var testQueries *domain.Queries

// func TestMain(m *testing.M) {
// 	db, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		panic(err)
// 	}

// 	testQueries = domain.New(db)

// 	m.Run()
// }

func TestCreateAccount(t *testing.T) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	testQueries = domain.New(db)

	arg := domain.CreateAccountParams{
		Owner:    "test",
		Balance:  100,
		Currency: "USD",
	}

	acount, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acount)

	require.Equal(t, arg.Owner, acount.Owner)
	require.Equal(t, arg.Balance, acount.Balance)
	require.Equal(t, arg.Currency, acount.Currency)
	require.NotZero(t, acount.ID)
}
