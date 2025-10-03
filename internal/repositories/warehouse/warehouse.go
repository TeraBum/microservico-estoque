package warehouse

import (
	"api-estoque/internal/config"
	warehouse "api-estoque/internal/model/warehouse"
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

func (r *Repository) List() (*[]warehouse.Warehouse, error) {
	ctx := context.Background()

	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "Name", "Location", "CreatedAt"
		FROM "Warehouse"
		ORDER BY "CreatedAt" DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []warehouse.Warehouse
	for rows.Next() {
		var w warehouse.Warehouse
		if err := rows.Scan(
			&w.Id,
			&w.Name,
			&w.Location,
			&w.CreatedAt,
		); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return &warehouses, nil
}

func (r *Repository) Create(w *warehouse.Warehouse) (*warehouse.Warehouse, error) {
	ctx := context.Background()
	query := `
		INSERT INTO "Warehouse" ("Name", "Location")
		VALUES ($1, $2)
		RETURNING "Id", "CreatedAt"
	`
	err := r.DB.QueryRow(ctx, query,
		w.Name,
		w.Location,
	).Scan(&w.Id, &w.CreatedAt)

	if err != nil {
		return nil, err
	}
	return w, nil
}

func (r *Repository) GetByID(id *uuid.UUID) (*warehouse.Warehouse, error) {
	ctx := context.Background()
	query := `
		SELECT "Id", "Name", "Location", "CreatedAt"
		FROM "Warehouse"
		WHERE "Id"=$1
	`
	var w warehouse.Warehouse
	err := r.DB.QueryRow(ctx, query, *id).Scan(
		&w.Id,
		&w.Name,
		&w.Location,
		&w.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *Repository) Update(w *warehouse.Warehouse) error {
	ctx := context.Background()

	query := `
		UPDATE "Warehouse"
		SET "Name"=$1, "Location"=$2
		WHERE "Id"=$3
		RETURNING "CreatedAt"
	`
	return r.DB.QueryRow(ctx, query,
		w.Name,
		w.Location,
		w.Id,
	).Scan(&w.CreatedAt)
}

func (r *Repository) Delete(id *uuid.UUID) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
		DELETE FROM "Warehouse"
		WHERE "Id"=$1
	`, *id)

	if err != nil {
		return fmt.Errorf("delete warehouse: %w", err)
	}
	return nil
}
