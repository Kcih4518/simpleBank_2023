package db

import (
	"context"
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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := New(tx)
	err = fn(queries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

// Implement money transfer transaction

// TransferTxParams contains the input parameters for the TransferTx function.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult contains the result of the TransferTx function.
// The created Transfer record.
// The FromAccount after its balance is subtracted.
// The ToAccount after its its balance is added.
// The FromEntry of the account which records that money is moving out of the FromAccount.
// And the ToEntry of the account which records that money is moving in to the ToAccount.
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// We create a transfer record with amount equals to 10.
// We create an entry record for account 1 with amount equals to -10, since money is moving out of this account.
// We create another entry record for account 2, but with amount equals to 10, because money is moving in to this account.
// Then we update the balance of account 1 by subtracting 10 from it.
// And finally we update the balance of account 2 by adding 10 to it.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// should convert arg (type TransferTxParams) to CreateTransferParams instead of using struct literal (gosimple)
		// Because CreateTransferParams has same fields as TransferTxParams.
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}
		return nil

		// TODO: Before implement update accounts' balance we have to solve db lock problem.
	})

	return result, err
}
