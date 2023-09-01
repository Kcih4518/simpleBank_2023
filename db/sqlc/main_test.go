package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Kcih4518/simpleBank_2023/util"
	_ "github.com/lib/pq"
)

// testQueries is a global variable that will be initialized once
// and can be used across multiple tests in the package.
// NewStore function requires *Queries and *sql.DB, so we need to export both.
var (
	testQueries *Queries
	testDB      *sql.DB
)

// TestMain is a special function recognized by the Go testing tool.
// It acts as an entry point to all the tests in the package.
// The setup and teardown for the entire test suite can be done here.
func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Establish a database connection.
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		// If there's an error, log it and exit.
		log.Fatal("cannot connect to db: ", err)
	}

	// Initialize testQueries with the database connection,
	// leveraging functions from the db.go file, even though it doesn't have a _test.go suffix.
	testQueries = New(testDB)

	// m.Run() executes all the tests in this package.
	// We capture the exit code to pass it to os.Exit later.
	exitCode := m.Run()

	// Close the database connection if needed.

	// Exit with the captured exit code. This is important to return a non-zero exit code
	// if any of the tests failed.
	os.Exit(exitCode)
}
