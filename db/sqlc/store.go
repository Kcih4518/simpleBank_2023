package db

import (
	"database/sql"
)

// Introduced `Store` struct, encapsulating `Queries` and integrating with `sql.DB` for transactional operations.
// Implemented `NewStore` function for initializing `Store`.
// Added `execTx` method to support customized operations within a transaction, automatically handling commit and rollback.

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	queries := New(tx)
// 	err = fn(queries)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return rbErr
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
