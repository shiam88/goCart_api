package postgres

import (
	"context"
	"database/sql"
)

type TransactionManager struct {
	db *DB
}

func NewTransactionManager(db *DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (m *TransactionManager) WithinTransaction(ctx context.Context,
	fn func(context.Context) error) error {
	if txFromContext(ctx) != nil {
		return fn(ctx)
	}

	tx, err := m.db.raw.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	},
	)

	if err != nil {
		return err
	}

	txctx := withTx(ctx, tx)

	if err := fn(txctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
