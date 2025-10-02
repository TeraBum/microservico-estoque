package product

import (
	"api-estoque/internal/config"
	productModel "api-estoque/internal/model/product"
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
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

func (r *Repository) List() (*[]productModel.Product, error) {
	ctx := context.Background()

	rows, err := r.DB.Query(ctx, `
		SELECT "Id", "CreatedAt", "Name", "Description", "Price", "Category", "ImagesJson", "IsActive"
		FROM "Product"
		ORDER BY "CreatedAt" DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []productModel.Product
	for rows.Next() {
		var p productModel.Product
		if err := rows.Scan(
			&p.Id,
			&p.CreatedAt,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Category,
			&p.ImagesJson,
			&p.IsActive,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &products, nil
}

func (r *Repository) Create(p *productModel.Product) (*productModel.Product, error) {
	ctx := context.Background()

	query := `
		INSERT INTO "Product" ("Name", "Description", "Price", "Category", "ImagesJson", "IsActive")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "Id", "CreatedAt"
	`
	err := r.DB.QueryRow(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Category,
		p.ImagesJson,
		p.IsActive,
	).Scan(&p.Id, &p.CreatedAt)

	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Repository) GetByID(id *uuid.UUID) (*productModel.Product, error) {
	ctx := context.Background()

	query := `
		SELECT "Id", "CreatedAt", "Name", "Description", "Price", "Category", "ImagesJson", "IsActive"
		FROM "Product"
		WHERE "Id" = $1
	`
	var p productModel.Product
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&p.Id,
		&p.CreatedAt,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Category,
		&p.ImagesJson,
		&p.IsActive,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(p *productModel.Product) error {
	ctx := context.Background()

	query := `
		UPDATE "Product"
		SET "Name"=$1, "Description"=$2, "Price"=$3, "Category"=$4, "ImagesJson"=$5, "IsActive"=$6
		WHERE "Id"=$7
		RETURNING "CreatedAt"
	`
	_, err := r.DB.Exec(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Category,
		p.ImagesJson,
		p.IsActive,
		p.Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id *uuid.UUID) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `DELETE FROM "Product" WHERE "Id"=$1`, id)
	if err != nil {
		return fmt.Errorf("delete product: %w", err)
	}
	return nil
}
