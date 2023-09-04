package domain

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

func TestMain(m *testing.M) {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	testQueries = New(testDB)

	m.Run()
}

func createRandomString() string {
	var randomString string
	randomString = time.Now().String()
	// get last 3 number
	randomString = randomString[len(randomString)-3:]
	return randomString
}

func createRandomNumber() int32 {
	randomNumber := int32(time.Now().UnixNano())
	return randomNumber
}
