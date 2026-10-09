package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	raw *sql.DB
}

func Open(ctx context.Context, dsn string) (*DB, error) {
	raw, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	raw.SetMaxOpenConns(20)
	raw.SetMaxIdleConns(10)
	raw.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	if err := raw.PingContext(pingCtx); err != nil {
		raw.Close()
		return nil, fmt.Errorf("postgres ping: %w", err)
	}

	return &DB{
		raw: raw,
	}, nil
}

func (d *DB) Close() error {
	return d.raw.Close()
}

func (d *DB) Raw() *sql.DB {
	return d.raw
}

type Execer interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type txKey struct{}

func withTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func txFromContext(ctx context.Context) *sql.Tx {
	tx, _ := ctx.Value(txKey{}).(*sql.Tx)
	return tx
}

func InTransaction(ctx context.Context) bool {
	return txFromContext(ctx) != nil
}

func (d *DB) Executor(ctx context.Context) Execer {
	if tx := txFromContext(ctx); tx != nil {
		return tx
	}
	return d.raw
}
