package stockitems

import (
	"api-estoque/internal/config"
	stockitems "api-estoque/internal/model/stock_items"
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

	setParts := []string{}
	args := []any{}
	argPos := 1

	if s.Quantity != nil {
		setParts = append(setParts, `"Quantity"=$`+strconv.Itoa(argPos))
		args = append(args, *s.Quantity)
		argPos++
	}

	if s.Reserved != nil {
		setParts = append(setParts, `"Reserved"=$`+strconv.Itoa(argPos))
		args = append(args, *s.Reserved)
		argPos++
	}

	if len(setParts) > 0 {
		setParts = append(setParts, `"UpdatedAt"=now()`)
	} else {
		return nil
	}

	query := `
		UPDATE "StockItems"
		SET ` + strings.Join(setParts, ", ") + `
		WHERE "ProductId"=$` + strconv.Itoa(argPos) + ` AND "WarehouseId"=$` + strconv.Itoa(argPos+1) + `
		RETURNING "UpdatedAt"
	`

	args = append(args, s.ProductId, s.WarehouseId)

	return r.DB.QueryRow(ctx, query, args...).Scan(&s.UpdatedAt)
}

func (r *Repository) DeductQuantity(baixa *stockitems.StockItemsBaixa) error {
	ctx := context.Background()

	query := `
		UPDATE "StockItems"
		SET "Quantity" = "Quantity" - $1,
		    "UpdatedAt" = now()
		WHERE "WarehouseId" = $2 
		  AND "ProductId" = $3
		  AND "Quantity" >= $1
		RETURNING "Quantity"
	`

	var newQuantity int
	err := r.DB.QueryRow(ctx, query, *baixa.Quantity, *baixa.WarehouseId, *baixa.ProductId).Scan(&newQuantity)
	if err != nil {
		return fmt.Errorf("failed to deduct quantity: %w", err)
	}

	if newQuantity < 0 {
		return fmt.Errorf("quantity cannot be negative")
	}

	return nil
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
