package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository { return &OrderRepository{db: db} }
func (r *OrderRepository) Save(ctx context.Context, o *order.Order) error {
	if postgres.InTransaction(ctx) {
		return r.saveWith(ctx, r.db.Executor(ctx), o)
	}
	tx, err := r.db.Raw().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := r.saveWith(ctx, tx, o); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *OrderRepository) saveWith(ctx context.Context, ex postgres.Execer, o *order.Order) error {
	_, err := ex.ExecContext(ctx, `INSERT INTO orders(id,user_id,status,subtotal,currency,shipping_fee,discount,total,shipping_address,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) ON CONFLICT(id) DO UPDATE SET status=EXCLUDED.status,subtotal=EXCLUDED.subtotal,shipping_fee=EXCLUDED.shipping_fee,discount=EXCLUDED.discount,total=EXCLUDED.total,shipping_address=EXCLUDED.shipping_address,updated_at=EXCLUDED.updated_at`, o.ID(), o.UserID(), string(o.Status()), o.Subtotal().Amount(), o.Total().Currency(), o.ShippingFee().Amount(), o.Discount().Amount(), o.Total().Amount(), o.ShippingAddress(), o.CreatedAt(), o.UpdatedAt())
	if err != nil {
		return err
	}
	_, err = ex.ExecContext(ctx, `DELETE FROM order_items WHERE order_id=$1`, o.ID())
	if err != nil {
		return err
	}
	for _, it := range o.Items() {
		sub, err := it.Subtotal()
		if err != nil {
			return err
		}
		_, err = ex.ExecContext(ctx, `INSERT INTO order_items(order_id,product_id,product_name,unit_price_amount,currency,quantity,subtotal) VALUES($1,$2,$3,$4,$5,$6,$7)`, o.ID(), it.ProductID(), it.ProductName(), it.UnitPrice().Amount(), it.UnitPrice().Currency(), it.Quantity(), sub.Amount())
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *OrderRepository) FindByID(ctx context.Context, id shared.ID) (*order.Order, error) {
	return r.findOne(ctx, `SELECT id,user_id,status,subtotal,currency,shipping_fee,discount,total,shipping_address,created_at,updated_at FROM orders WHERE id=$1`, id)
}
func (r *OrderRepository) findOne(ctx context.Context, q string, arg any) (*order.Order, error) {
	var id, uid shared.ID
	var status, currency, address string
	var subtotal, shipping, discount, total int64
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, q, arg).Scan(&id, &uid, &status, &subtotal, &currency, &shipping, &discount, &total, &address, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, order.ErrOrderNotFound
	}
	if err != nil {
		return nil, err
	}
	items, err := r.loadItems(ctx, id)
	if err != nil {
		return nil, err
	}
	sm, _ := shared.NewMoney(subtotal, currency)
	sf, _ := shared.NewMoney(shipping, currency)
	dm, _ := shared.NewMoney(discount, currency)
	tm, _ := shared.NewMoney(total, currency)
	return order.Restore(id, uid, items, order.Status(status), sm, sf, dm, tm, address, created, updated)
}
func (r *OrderRepository) loadItems(ctx context.Context, orderID shared.ID) ([]order.Item, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT product_id,product_name,unit_price_amount,currency,quantity FROM order_items WHERE order_id=$1 ORDER BY id`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []order.Item{}
	for rows.Next() {
		var pid shared.ID
		var name, currency string
		var amount int64
		var qty int
		if err := rows.Scan(&pid, &name, &amount, &currency, &qty); err != nil {
			return nil, err
		}
		m, err := shared.NewMoney(amount, currency)
		if err != nil {
			return nil, err
		}
		it, err := order.NewItem(pid, name, m, qty)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}
func (r *OrderRepository) ListByUserID(ctx context.Context, userID shared.ID, p shared.Pagination) ([]*order.Order, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT id FROM orders WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userID, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids := []shared.ID{}
	for rows.Next() {
		var id shared.ID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	out := []*order.Order{}
	for _, id := range ids {
		o, err := r.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, nil
}
func (r *OrderRepository) List(ctx context.Context, p shared.Pagination) ([]*order.Order, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT id FROM orders ORDER BY created_at DESC LIMIT $1 OFFSET $2`, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids := []shared.ID{}
	for rows.Next() {
		var id shared.ID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	out := []*order.Order{}
	for _, id := range ids {
		o, err := r.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, rows.Err()
}
func (r *OrderRepository) CountByUserID(ctx context.Context, userID shared.ID) (int, error) {
	var n int
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT COUNT(*) FROM orders WHERE user_id=$1`, userID).Scan(&n)
	return n, err
}
