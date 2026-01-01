package repository

import (
	"database/sql"
	// "errors" 
	"vending-machine/internal/migrations"
	

)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrder(tx *sql.Tx, order migrations.Orders) error {
	_, err := tx.Exec(`
		INSERT INTO orders (reference, inventory_id, product_variant_id, quantity, email)
		VALUES (?, ?, ?, ?, ?)
	`, order.InvoiceRef, order.InventoryID, order.ProductVariantID, order.Quantity, order.Email)

	return err
}



func (r *OrderRepository) UpdateStatus(ref, status string) error {
    _, err := r.db.Exec(`
        UPDATE orders SET status = ?
        WHERE reference = ?
    `, status, ref)
    return err
}