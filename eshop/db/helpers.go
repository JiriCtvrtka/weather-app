package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestTable(t *testing.T, db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test (id varchar(255) PRIMARY KEY)")
	require.NoError(t, err)
}
