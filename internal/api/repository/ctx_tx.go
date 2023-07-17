package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// ExtractTx extracts transaction from context
func ExtractTx(ctx context.Context) *pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(*pgx.Tx); ok {
		return tx
	}
	return nil
}

// TX is an interface for working with transactions.
type TX interface {
	WithTx(context.Context, func(context.Context) error) error
}

// txRepo is a base repository for working with transactions.
type txRepo struct {
	db *pgx.Conn
}

// NewTxRepo constructor.
func NewTxRepo(db *pgx.Conn) *txRepo {
	return &txRepo{
		db: db,
	}
}

// WithTx executes fn with transaction.
func (r *txRepo) WithTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fn(injectTx(ctx, &tx)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
