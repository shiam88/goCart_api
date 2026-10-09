package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
	"goCartAPI/infrastructure/postgres"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `INSERT INTO users (id,name,email,password_hash,role,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)
	ON CONFLICT (id) DO UPDATE SET name=EXCLUDED.name,email=EXCLUDED.email,password_hash=EXCLUDED.password_hash,role=EXCLUDED.role,updated_at=EXCLUDED.updated_at`,
		u.ID(), u.Name(), u.Email(), u.PasswordHash(), string(u.Role()), u.CreatedAt(), u.UpdatedAt())
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id shared.ID) (*user.User, error) {
	return r.find(ctx, `SELECT id,name,email,password_hash,role,created_at,updated_at FROM users WHERE id=$1`, id)
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return r.find(ctx, `SELECT id,name,email,password_hash,role,created_at,updated_at FROM users WHERE lower(email)=lower($1)`, strings.TrimSpace(email))
}
func (r *UserRepository) find(ctx context.Context, query string, arg any) (*user.User, error) {
	var id shared.ID
	var name, email, hash, role string
	var created, updated time.Time
	err := r.db.Executor(ctx).QueryRowContext(ctx, query, arg).Scan(&id, &name, &email, &hash, &role, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, user.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user.Restore(id, name, email, hash, user.Role(role), created, updated)
}
func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.Executor(ctx).QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE lower(email)=lower($1))`, strings.TrimSpace(email)).Scan(&exists)
	return exists, err
}
