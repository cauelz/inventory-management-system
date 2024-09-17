package models

import "time"

type Product struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	StockQty    int       `db:"stock_quantity" json:"stock_quantity"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
