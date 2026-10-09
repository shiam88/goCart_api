package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"goCartAPI/domain/product"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type ProductRepository struct{ db *postgres.DB }

func NewProductRepository(db *postgres.DB) *ProductRepository { return &ProductRepository{db: db} }

func (r *ProductRepository) Save(ctx context.Context, p *product.Product) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `INSERT INTO products(id,category_id,name,slug,sku,description,price_amount,currency,image_url,active,created_at,updated_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	ON CONFLICT(id) DO UPDATE SET category_id=EXCLUDED.category_id,name=EXCLUDED.name,slug=EXCLUDED.slug,sku=EXCLUDED.sku,description=EXCLUDED.description,price_amount=EXCLUDED.price_amount,currency=EXCLUDED.currency,image_url=EXCLUDED.image_url,active=EXCLUDED.active,updated_at=EXCLUDED.updated_at`,
		p.ID(), p.CategoryID(), p.Name(), p.Slug(), p.SKU(), p.Description(), p.Price().Amount(), p.Price().Currency(), p.ImageURL(), p.IsActive(), p.CreatedAt(), p.UpdatedAt())
	return err
}

func (r *ProductRepository) FindByID(ctx context.Context, id shared.ID) (*product.Product, error) {
	return r.find(ctx, `SELECT id,category_id,name,slug,sku,description,price_amount,currency,image_url,active,created_at,updated_at FROM products WHERE id=$1`, id)
}

func (r *ProductRepository) FindBySKU(ctx context.Context, sku string) (*product.Product, error) {
	return r.find(ctx, `SELECT id,category_id,name,slug,sku,description,price_amount,currency,image_url,active,created_at,updated_at FROM products WHERE upper(sku)=upper($1)`, strings.TrimSpace(sku))
}

func (r *ProductRepository) find(ctx context.Context, q string, arg any) (*product.Product, error) {
	var id, catID shared.ID
	var name, slug, sku, desc, currency, image string
	var amount int64
	var active bool
	var created, updated time.Time
	if err := r.db.Executor(ctx).QueryRowContext(ctx, q, arg).Scan(&id, &catID, &name, &slug, &sku, &desc, &amount, &currency, &image, &active, &created, &updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, product.ErrProductNotFound
		}
		return nil, err
	}
	money, err := shared.NewMoney(amount, currency)
	if err != nil {
		return nil, err
	}
	return product.Restore(id, catID, name, slug, sku, desc, money, image, active, created, updated)
}

func productWhere(f product.Filter) (string, []any) {
	parts := []string{
		"1=1",
	}
	args := []any{}
	if search := strings.TrimSpace(f.Search); search != "" {
		args = append(args, search)
		pos := len(args)
		parts = append(parts, fmt.Sprintf("(name ILIKE '%%' || $%d || '%%' OR sku ILIKE '%%' || $%d || '%%')", pos, pos))
	}
	if f.CategoryID != nil {
		args = append(args, *f.CategoryID)
		parts = append(parts, fmt.Sprintf("category_id=$%d", len(args)))
	}
	if f.MinPrice != nil {
		args = append(args, *f.MinPrice)
		parts = append(parts, fmt.Sprintf("price_amount >= $%d", len(args)))
	}
	if f.MaxPrice != nil {
		args = append(args, *f.MaxPrice)
		parts = append(parts, fmt.Sprintf("price_amount <= $%d", len(args)))
	}
	if f.OnlyActive {
		parts = append(parts, "active=true")
	}
	return strings.Join(parts, " AND "), args
}

func (r *ProductRepository) List(ctx context.Context, f product.Filter, p shared.Pagination) ([]*product.Product, error) {
	where, args := productWhere(f)
	args = append(args, p.Limit)
	limitPos := len(args)
	args = append(args, p.Offset())
	offsetPos := len(args)
	q := fmt.Sprintf(`SELECT id,category_id,name,slug,sku,description,price_amount,currency,image_url,active,created_at,updated_at FROM products WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, where, limitPos, offsetPos)
	rows, err := r.db.Executor(ctx).QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []*product.Product{}
	for rows.Next() {
		var id, catID shared.ID
		var name, slug, sku, desc, currency, image string
		var amount int64
		var active bool
		var created, updated time.Time
		if err := rows.Scan(&id, &catID, &name, &slug, &sku, &desc, &amount, &currency, &image, &active, &created, &updated); err != nil {
			return nil, err
		}
		money, err := shared.NewMoney(amount, currency)
		if err != nil {
			return nil, err
		}
		p, err := product.Restore(id, catID, name, slug, sku, desc, money, image, active, created, updated)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (r *ProductRepository) Count(ctx context.Context, f product.Filter) (int, error) {
	where, args := productWhere(f)
	var n int
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT COUNT(*) FROM products WHERE `+where, args...).Scan(&n)
	return n, err
}

func (r *ProductRepository) Delete(ctx context.Context, id shared.ID) error {
	res, err := r.db.Executor(ctx).ExecContext(ctx, `DELETE FROM products WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return product.ErrProductNotFound
	}
	return nil
}

func (r *ProductRepository) SKUExists(ctx context.Context, sku string) (bool, error) {
	var exists bool
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM products WHERE upper(sku)=upper($1))`, strings.TrimSpace(sku)).Scan(&exists)
	return exists, err
}
