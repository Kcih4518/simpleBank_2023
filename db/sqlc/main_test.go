package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// testQueries is a global variable that will be initialized once
// and can be used across multiple tests in the package.
var testQueries *Queries

// TestMain is a special function recognized by the Go testing tool.
// It acts as an entry point to all the tests in the package.
// The setup and teardown for the entire test suite can be done here.
func TestMain(m *testing.M) {
	// Establish a database connection.
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		// If there's an error, log it and exit.
		log.Fatal("cannot connect to db: ", err)
	}

	// Initialize testQueries with the database connection,
	// leveraging functions from the db.go file, even though it doesn't have a _test.go suffix.
	testQueries = New(conn)

	// m.Run() executes all the tests in this package.
	// We capture the exit code to pass it to os.Exit later.
	exitCode := m.Run()

	// Close the database connection if needed.

	// Exit with the captured exit code. This is important to return a non-zero exit code
	// if any of the tests failed.
	os.Exit(exitCode)
}
