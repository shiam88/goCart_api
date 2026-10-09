package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type CartRepository struct{ db *postgres.DB }

func NewCartRepository(db *postgres.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) Save(ctx context.Context, c *cart.Cart) error {
	if postgres.InTransaction(ctx) {
		return r.saveWith(ctx, r.db.Executor(ctx), c)
	}
	tx, err := r.db.Raw().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := r.saveWith(ctx, tx, c); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *CartRepository) saveWith(ctx context.Context, ex postgres.Execer, c *cart.Cart) error {
	_, err := ex.ExecContext(ctx, `INSERT INTO carts(id,user_id,currency,created_at,updated_at) VALUES($1,$2,$3,$4,$5) ON CONFLICT(user_id) DO UPDATE SET currency=EXCLUDED.currency,updated_at=EXCLUDED.updated_at`, c.ID(), c.UserID(), c.Currency(), c.CreatedAt(), c.UpdatedAt())
	if err != nil {
		return err
	}
	_, err = ex.ExecContext(ctx, `DELETE FROM cart_items WHERE cart_id=$1`, c.ID())
	if err != nil {
		return err
	}
	for _, it := range c.Items() {
		_, err = ex.ExecContext(ctx, `INSERT INTO cart_items(cart_id,product_id,product_name,unit_price_amount,currency,quantity) VALUES($1,$2,$3,$4,$5,$6)`, c.ID(), it.ProductID(), it.ProductName(), it.UnitPrice().Amount(), it.UnitPrice().Currency(), it.Quantity())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CartRepository) FindByID(ctx context.Context, id shared.ID) (*cart.Cart, error) {
	var cartID, userID shared.ID
	var currency string
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT id,user_id,currency,created_at,updated_at FROM carts WHERE id=$1`, id).Scan(&cartID, &userID, &currency, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, cart.ErrCartNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.restoreCart(ctx, cartID, userID, currency, created, updated)
}

func (r *CartRepository) restoreCart(ctx context.Context, id, userID shared.ID, currency string, created, updated time.Time) (*cart.Cart, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT product_id,product_name,unit_price_amount,currency,quantity FROM cart_items WHERE cart_id=$1 ORDER BY id`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*cart.Item{}
	for rows.Next() {
		var pid shared.ID
		var name, cur string
		var amount int64
		var qty int
		if err := rows.Scan(&pid, &name, &amount, &cur, &qty); err != nil {
			return nil, err
		}
		m, err := shared.NewMoney(amount, cur)
		if err != nil {
			return nil, err
		}
		it, err := cart.NewItem(pid, name, m, qty)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cart.Restore(id, userID, currency, items, created, updated)
}

func (r *CartRepository) FindByUserID(ctx context.Context, userID shared.ID) (*cart.Cart, error) {
	var id, uid shared.ID
	var currency string
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT id,user_id,currency,created_at,updated_at FROM carts WHERE user_id=$1`, userID).Scan(&id, &uid, &currency, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, cart.ErrCartNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.restoreCart(ctx, id, uid, currency, created, updated)
}

func (r *CartRepository) Delete(ctx context.Context, id shared.ID) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `DELETE FROM carts WHERE id=$1`, id)
	return err
}
