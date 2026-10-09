package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type PaymentRepository struct{ db *postgres.DB }

func NewPaymentRepository(db *postgres.DB) *PaymentRepository { return &PaymentRepository{db: db} }
func (r *PaymentRepository) Save(ctx context.Context, p *payment.Transaction) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `INSERT INTO payment_transactions(id,order_id,user_id,amount,currency,provider,status,provider_reference,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) ON CONFLICT(id) DO UPDATE SET status=EXCLUDED.status,provider_reference=EXCLUDED.provider_reference,updated_at=EXCLUDED.updated_at`, p.ID(), p.OrderID(), p.UserID(), p.Amount().Amount(), p.Amount().Currency(), string(p.Provider()), string(p.Status()), p.ProviderReference(), p.CreatedAt(), p.UpdatedAt())
	return err
}
func (r *PaymentRepository) FindByID(ctx context.Context, id shared.ID) (*payment.Transaction, error) {
	return r.find(ctx, `SELECT id,order_id,user_id,amount,currency,provider,status,provider_reference,created_at,updated_at FROM payment_transactions WHERE id=$1`, id)
}
func (r *PaymentRepository) FindByOrderID(ctx context.Context, orderID shared.ID) (*payment.Transaction, error) {
	return r.find(ctx, `SELECT id,order_id,user_id,amount,currency,provider,status,provider_reference,created_at,updated_at FROM payment_transactions WHERE order_id=$1 ORDER BY created_at DESC LIMIT 1`, orderID)
}
func (r *PaymentRepository) find(ctx context.Context, q string, arg any) (*payment.Transaction, error) {
	var id, oid, uid shared.ID
	var amount int64
	var currency, providerName, status, ref string
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, q, arg).Scan(&id, &oid, &uid, &amount, &currency, &providerName, &status, &ref, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, payment.ErrTransactionNotFound
	}
	if err != nil {
		return nil, err
	}
	m, err := shared.NewMoney(amount, currency)
	if err != nil {
		return nil, err
	}
	return payment.Restore(id, oid, uid, m, payment.Provider(providerName), payment.Status(status), ref, created, updated)
}
func (r *PaymentRepository) ListByUserID(ctx context.Context, userID shared.ID, p shared.Pagination) ([]*payment.Transaction, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT id,order_id,user_id,amount,currency,provider,status,provider_reference,created_at,updated_at FROM payment_transactions WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userID, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []*payment.Transaction{}
	for rows.Next() {
		var id, oid, uid shared.ID
		var amount int64
		var cur, prov, st, ref string
		var created, updated time.Time
		if err := rows.Scan(&id, &oid, &uid, &amount, &cur, &prov, &st, &ref, &created, &updated); err != nil {
			return nil, err
		}
		m, err := shared.NewMoney(amount, cur)
		if err != nil {
			return nil, err
		}
		t, err := payment.Restore(id, oid, uid, m, payment.Provider(prov), payment.Status(st), ref, created, updated)
		if err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
