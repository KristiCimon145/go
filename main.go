package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// SafeQueryContext ensures that if a query context is canceled,
// the transaction is rolled back to prevent connection leaks.
func SafeQueryContext(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			// Ensure the transaction is closed if the query context was canceled
			_ = tx.Rollback()
		}
		return nil, err
	}
	return rows, nil
}

func main() {
	fmt.Println("Use SafeQueryContext to prevent connection leaks during context cancellation.")
}