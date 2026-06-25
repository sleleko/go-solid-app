package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"go-order-solid/internal/domain"
)

// SQLiteRepo — конкретная реализация репозитория для SQLite.
// Ее можно заменить на PostgreSQLRepo, MySQLRepo и т.д.,
// если новая структура реализует нужные интерфейсы.
type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) Init(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer TEXT NOT NULL,
			products TEXT NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL
		)`)
	if err != nil {
		return fmt.Errorf("init sqlite repository: %w", err)
	}

	return nil
}

func (r *SQLiteRepo) SaveOrder(ctx context.Context, order domain.Order) error {
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return fmt.Errorf("marshal products: %w", err)
	}

	_, err = r.db.ExecContext(
		ctx,
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer,
		string(productsJSON),
		order.Total,
		order.Status,
	)
	if err != nil {
		return fmt.Errorf("save order: %w", err)
	}

	return nil
}
