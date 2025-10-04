package product

import (
	"api-estoque/internal/config"
	productModel "api-estoque/internal/model/product"
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

func (r *Repository) Create(p *productModel.Product) (*uuid.UUID, error) {
	ctx := context.Background()

	query := `
		INSERT INTO "Product" ("Name", "Description", "Price", "Category", "ImagesJson", "IsActive")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "Id"
	`
	err := r.DB.QueryRow(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Category,
		p.ImagesJson,
		p.IsActive,
	).Scan(&p.Id)

	if err != nil {
		return nil, err
	}
	return p.Id, nil
}

func (r *Repository) GetByID(id *uuid.UUID) (*productModel.Product, error) {
	ctx := context.Background()
	query := `
		SELECT "Id", "CreatedAt", "Name", "Description", "Price", "Category", "ImagesJson", "IsActive"
		FROM "Product"
		WHERE "Id"=$1
	`

	var p productModel.Product
	err := r.DB.QueryRow(ctx, query, *id).Scan(
		&p.Id,
		&p.CreatedAt,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Category,
		&p.ImagesJson,
		&p.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(p *productModel.Product) error {
	ctx := context.Background()

	setParts := []string{}
	args := []any{}
	argPos := 1

	if p.Name != nil {
		setParts = append(setParts, `"Name"=$`+strconv.Itoa(argPos))
		args = append(args, *p.Name)
		argPos++
	}

	if p.Description != nil {
		setParts = append(setParts, `"Description"=$`+strconv.Itoa(argPos))
		args = append(args, *p.Description)
		argPos++
	}

	if p.Price != nil {
		setParts = append(setParts, `"Price"=$`+strconv.Itoa(argPos))
		args = append(args, *p.Price)
		argPos++
	}

	if p.Category != nil {
		setParts = append(setParts, `"Category"=$`+strconv.Itoa(argPos))
		args = append(args, *p.Category)
		argPos++
	}

	if p.ImagesJson != nil {
		setParts = append(setParts, `"ImagesJson"=$`+strconv.Itoa(argPos))
		args = append(args, *p.ImagesJson)
		argPos++
	}

	if p.IsActive != nil {
		setParts = append(setParts, `"IsActive"=$`+strconv.Itoa(argPos))
		args = append(args, *p.IsActive)
		argPos++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := `
		UPDATE "Product"
		SET ` + strings.Join(setParts, ", ") + `
		WHERE "Id"=$` + strconv.Itoa(argPos)
	args = append(args, p.Id)

	_, err := r.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("update product: %w", err)
	}

	return nil
}

func (r *Repository) Delete(id *uuid.UUID) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx, `
		DELETE FROM "Product"
		WHERE "Id"=$1
	`, *id)

	if err != nil {
		return fmt.Errorf("delete product: %w", err)
	}
	return nil
}
