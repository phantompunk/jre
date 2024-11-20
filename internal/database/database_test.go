package sql

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

// CreateDatabase for testing.
func CreateDatabase(t *testing.T) *Database {
	t.Helper()

	db := NewDatabase(WithBaseUrl("../database/db.db"))
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestGetRandomQuote(t *testing.T) {
	t.Run("Get a Random Quote", func(t *testing.T) {
		db := CreateDatabase(t)
		require.NotNil(t, db)
		quote, err := db.GetRandomQoute(context.Background())
		require.NoError(t, err)
		require.NotEmpty(t, quote.Quote)
	})
}

func TestGetQuoteByID(t *testing.T) {
	t.Run("Get Quote By ID", func(t *testing.T) {
		db := CreateDatabase(t)
		require.NotNil(t, db)
		id := "ba28b5"
		quote, err := db.GetQouteById(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, id, quote.ShortID)
		require.NotEmpty(t, quote.Quote)
	})
}
