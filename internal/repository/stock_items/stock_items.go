package stockitems

import (
	"api-estoque/internal/config"
	stockitems "api-estoque/internal/model/stock_items"
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func New() *Repository {
	maxConns := 10
	maxIdleTime := 30 * time.Second
	maxLifetime := 2 * time.Minute

	return &Repository{
		DB: config.PostgresConn(maxConns, maxIdleTime, maxLifetime),
	}
}

func (r *Repository) List() (*[]stockitems.StockItems, error) {
	ctx := context.Background()

	rows, err := r.DB.Query(ctx, `
		SELECT "ProductId", "WarehouseId", "Quantity", "Reserved", "UpdatedAt"
		FROM "StockItems"
		ORDER BY "UpdatedAt" DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []stockitems.StockItems
	for rows.Next() {
		var s stockitems.StockItems
		if err := rows.Scan(
			&s.ProductId,
			&s.WarehouseId,
			&s.Quantity,
			&s.Reserved,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, s)
	}
	return &items, nil
}

func (r *Repository) Create(s *stockitems.StockItems) (*stockitems.StockItems, error) {
	ctx := context.Background()
	query := `
		INSERT INTO "StockItems" ("ProductId", "WarehouseId", "Quantity", "Reserved")
		VALUES ($1, $2, $3, $4)
		RETURNING "UpdatedAt"
	`
	err := r.DB.QueryRow(ctx, query,
		s.ProductId,
		s.WarehouseId,
		s.Quantity,
		s.Reserved,
	).Scan(&s.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) GetByID(idWarehouse *uuid.UUID, idProduct *uuid.UUID) (*stockitems.StockItems, error) {
	ctx := context.Background()
	query := `
		SELECT "ProductId", "WarehouseId", "Quantity", "Reserved", "UpdatedAt"
		FROM "StockItems"
		WHERE "WarehouseId"=$1 AND "ProductId"=$2
	`
	var s stockitems.StockItems
	err := r.DB.QueryRow(ctx, query, *idWarehouse, *idProduct).Scan(
		&s.ProductId,
		&s.WarehouseId,
		&s.Quantity,
		&s.Reserved,
		&s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) Update(s *stockitems.StockItems) error {
	ctx := context.Background()

	query := `
		UPDATE "StockItems"
		SET "Quantity"=$1, "Reserved"=$2, "UpdatedAt"=now()
		WHERE "ProductId"=$3 AND "WarehouseId"=$4
		RETURNING "UpdatedAt"
	`
	return r.DB.QueryRow(ctx, query,
		s.Quantity,
		s.Reserved,
		s.ProductId,
		s.WarehouseId,
	).Scan(&s.UpdatedAt)
}

func (r *Repository) Delete(idWarehouse *uuid.UUID, idProduct *uuid.UUID) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
		DELETE FROM "StockItems"
		WHERE "ProductId"=$1 AND "WarehouseId"=$2
	`, *idProduct, *idWarehouse)

	if err != nil {
		return fmt.Errorf("delete stock item: %w", err)
	}
	return nil
}
