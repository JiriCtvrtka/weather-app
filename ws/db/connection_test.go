package db

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	_, err := Connect()
	require.NoError(t, err)
}
