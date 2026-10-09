package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goCartAPI/domain/category"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type CategoryRepository struct {
	db *postgres.DB
}

func NewCategoryRepository(db *postgres.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Save(ctx context.Context, c *category.Category) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `INSERT INTO categories(id,name,slug,active,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6)
	ON CONFLICT(id) DO UPDATE SET name=EXCLUDED.name,slug=EXCLUDED.slug,active=EXCLUDED.active,updated_at=EXCLUDED.updated_at`, c.ID(), c.Name(), c.Slug(), c.IsActive(), c.CreatedAt(), c.UpdatedAt())
	return err
}

func (r *CategoryRepository) FindByID(ctx context.Context, id shared.ID) (*category.Category, error) {
	return r.find(
		ctx,
		`
		 SELECT
		  id,
		  name,
		  slug,
		  active,
		  created_at,
		  updated_at
		 FROM categories WHERE id=$1`, id,
	)
}
func (r *CategoryRepository) FindBySlug(ctx context.Context, slug string) (*category.Category, error) {
	return r.find(ctx, `SELECT id,name,slug,active,created_at,updated_at FROM categories WHERE slug=$1`, slug)
}

func (r *CategoryRepository) find(ctx context.Context, q string, arg any) (*category.Category, error) {

	var id shared.ID
	var name, slug string
	var active bool
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, q, arg).Scan(&id, &name, &slug, &active, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, category.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return category.Restore(id, name, slug, active, created, updated)
}

func (r *CategoryRepository) List(ctx context.Context, p shared.Pagination) ([]*category.Category, error) {
	rows, err := r.db.Executor(ctx).QueryContext(ctx, `SELECT id,name,slug,active,created_at,updated_at FROM categories ORDER BY name LIMIT $1 OFFSET $2`, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []*category.Category{}
	for rows.Next() {
		var id shared.ID
		var name, slug string
		var active bool
		var created, updated time.Time
		if err := rows.Scan(&id, &name, &slug, &active, &created, &updated); err != nil {
			return nil, err
		}
		c, err := category.Restore(id, name, slug, active, created, updated)
		if err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

/*
 */
func (r *CategoryRepository) Count(ctx context.Context) (int, error) {
	var n int
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT COUNT(*) FROM categories`).Scan(&n)
	return n, err
}
func (r *CategoryRepository) Delete(ctx context.Context, id shared.ID) error {
	res, err := r.db.Executor(ctx).ExecContext(ctx, `DELETE FROM categories WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return category.ErrCategoryNotFound
	}
	return nil
}
