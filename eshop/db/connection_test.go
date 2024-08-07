package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDBConnection(t *testing.T) {
	_, err := Connection()
	require.NoError(t, err)
}

func TestDB(t *testing.T) {
	db, err := Connection()
	require.NoError(t, err)
	defer db.Close()

	createTestTable(t, db)

	_, err = db.Query("SELECT * FROM test")
	require.NoError(t, err)

	_, err = db.Exec("DROP TABLE test")
	require.NoError(t, err)
}
