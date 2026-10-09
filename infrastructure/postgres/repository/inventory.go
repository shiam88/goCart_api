package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goCartAPI/domain/inventory"
	"goCartAPI/domain/shared"
	"goCartAPI/infrastructure/postgres"
)

type InventoryRepository struct{ db *postgres.DB }

func NewInventoryRepository(db *postgres.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}
func (r *InventoryRepository) Save(ctx context.Context, i *inventory.Inventory) error {
	_, err := r.db.Executor(ctx).ExecContext(ctx, `INSERT INTO inventory(product_id,on_hand,reserved,updated_at) VALUES($1,$2,$3,$4) ON CONFLICT(product_id) DO UPDATE SET on_hand=EXCLUDED.on_hand,reserved=EXCLUDED.reserved,updated_at=EXCLUDED.updated_at`, i.ProductID(), i.OnHand(), i.Reserved(), i.UpdatedAt())
	return err
}
func (r *InventoryRepository) FindByProductID(ctx context.Context, productID shared.ID) (*inventory.Inventory, error) {
	var id shared.ID
	var onHand, reserved int
	var updated time.Time
	q := `SELECT product_id,on_hand,reserved,updated_at FROM inventory WHERE product_id=$1`
	if postgres.InTransaction(ctx) {
		q += ` FOR UPDATE`
	}
	err := r.db.Executor(ctx).QueryRowContext(ctx, q, productID).Scan(&id, &onHand, &reserved, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, inventory.ErrInventoryNotFound
	}
	if err != nil {
		return nil, err
	}
	return inventory.Restore(id, onHand, reserved, updated)
}
