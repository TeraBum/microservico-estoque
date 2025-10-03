package stockmoves

import (
	"api-estoque/internal/config"
	stockmoves "api-estoque/internal/model/stock_moves"
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

// List returns all stock moves ordered by CreatedAt desc
func (r *Repository) List() (*[]stockmoves.StockMove, error) {
	ctx := context.Background()

	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "ProductId", "WarehouseId", "QtyMoved", "Reason", "CreatedAt"
		FROM "StockMoves"
		ORDER BY "CreatedAt" DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []stockmoves.StockMove
	for rows.Next() {
		var m stockmoves.StockMove
		if err := rows.Scan(
			&m.Id,
			&m.ProductId,
			&m.WarehouseId,
			&m.QtyMoved,
			&m.Reason,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		moves = append(moves, m)
	}
	return &moves, nil
}

// Create inserts a new stock move and returns it
func (r *Repository) Create(m *stockmoves.StockMove) (*stockmoves.StockMove, error) {
	ctx := context.Background()
	query := `
		INSERT INTO "StockMoves" ("ProductId", "WarehouseId", "QtyMoved", "Reason")
		VALUES ($1, $2, $3, $4)
		RETURNING "Id", "CreatedAt"
	`
	err := r.DB.QueryRow(ctx, query,
		m.ProductId,
		m.WarehouseId,
		m.QtyMoved,
		m.Reason,
	).Scan(&m.Id, &m.CreatedAt)

	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetByID fetches one stock move by its primary key
func (r *Repository) GetByID(id *uuid.UUID) (*stockmoves.StockMove, error) {
	ctx := context.Background()
	query := `
		SELECT "Id", "ProductId", "WarehouseId", "QtyMoved", "Reason", "CreatedAt"
		FROM "StockMoves"
		WHERE "Id"=$1
	`
	var m stockmoves.StockMove
	err := r.DB.QueryRow(ctx, query, *id).Scan(
		&m.Id,
		&m.ProductId,
		&m.WarehouseId,
		&m.QtyMoved,
		&m.Reason,
		&m.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ListByProduct fetches all moves for a given product
func (r *Repository) ListByProduct(productId *uuid.UUID) (*[]stockmoves.StockMove, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "ProductId", "WarehouseId", "QtyMoved", "Reason", "CreatedAt"
		FROM "StockMoves"
		WHERE "ProductId"=$1
		ORDER BY "CreatedAt" DESC
	`, *productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []stockmoves.StockMove
	for rows.Next() {
		var m stockmoves.StockMove
		if err := rows.Scan(
			&m.Id,
			&m.ProductId,
			&m.WarehouseId,
			&m.QtyMoved,
			&m.Reason,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		moves = append(moves, m)
	}
	return &moves, nil
}

func (r *Repository) ListByWarehouse(warehouseId *uuid.UUID) (*[]stockmoves.StockMove, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "ProductId", "WarehouseId", "QtyMoved", "Reason", "CreatedAt"
		FROM "StockMoves"
		WHERE "WarehouseId"=$1
		ORDER BY "CreatedAt" DESC
	`, *warehouseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []stockmoves.StockMove
	for rows.Next() {
		var m stockmoves.StockMove
		if err := rows.Scan(
			&m.Id,
			&m.ProductId,
			&m.WarehouseId,
			&m.QtyMoved,
			&m.Reason,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		moves = append(moves, m)
	}
	return &moves, nil
}

func (r *Repository) ListByWarehouseAndProduct(warehouseId *uuid.UUID, productId *uuid.UUID) (*[]stockmoves.StockMove, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "ProductId", "WarehouseId", "QtyMoved", "Reason", "CreatedAt"
		FROM "StockMoves"
		WHERE "WarehouseId"=$1 AND "ProductId"=$2
		ORDER BY "CreatedAt" DESC
	`, *warehouseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []stockmoves.StockMove
	for rows.Next() {
		var m stockmoves.StockMove
		if err := rows.Scan(
			&m.Id,
			&m.ProductId,
			&m.WarehouseId,
			&m.QtyMoved,
			&m.Reason,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		moves = append(moves, m)
	}
	return &moves, nil
}

// Delete removes a stock move by Id
func (r *Repository) Delete(id *uuid.UUID) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
		DELETE FROM "StockMoves"
		WHERE "Id"=$1
	`, *id)

	if err != nil {
		return fmt.Errorf("delete stock move: %w", err)
	}
	return nil
}
