package pg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect(connStr string) (*sql.DB, error) {

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

		return nil, fmt.Errorf("connect DB failed. Err: %w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "connect to DB successfully")

	return conn, nil
}
