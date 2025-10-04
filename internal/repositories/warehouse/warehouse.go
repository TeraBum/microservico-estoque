package warehouse

import (
	"api-estoque/internal/config"
	warehouse "api-estoque/internal/model/warehouse"
	"context"
	"fmt"
	"strconv"
	"strings"
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

	setParts := []string{}
	args := []any{}
	argPos := 1

	if w.Name != nil && *w.Name != "" {
		setParts = append(setParts, `"Name"=$`+strconv.Itoa(argPos))
		args = append(args, *w.Name)
		argPos++
	}

	if w.Location != nil && *w.Location != "" {
		setParts = append(setParts, `"Location"=$`+strconv.Itoa(argPos))
		args = append(args, *w.Location)
		argPos++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := `
		UPDATE "Warehouse"
		SET ` + strings.Join(setParts, ", ") + `
		WHERE "Id"=$` + strconv.Itoa(argPos) + `
		RETURNING "CreatedAt"
	`

	args = append(args, w.Id)

	return r.DB.QueryRow(ctx, query, args...).Scan(&w.CreatedAt)
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
