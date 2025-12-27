package repository

import (
	"database/sql"
	"errors"
	"vending-machine/internal/migrations"
)

type InventoryStockRepository struct {
	db *sql.DB
}

type AddStockItem struct  {
	ProductVariantID int64
	quantity int
}


func NewInventoryStockRepository(db *sql.DB) *InventoryStockRepository {
	return &InventoryStockRepository{db: db}
}

func (invs *InventoryStockRepository) AddMmultiStock(inventoryId int64, items []migrations.AddStockItem) error  {
	if len(items) == 0 {
		return nil
	}

	query := `
	INSERT INTO inventory_stock (inventory_id, product_variant_id, quantity)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
		quantity = quantity + VALUES(quantity)
	`

	txs, err := invs.db.Begin();
	if err != nil {
		txs.Rollback();
		return err
	}

	statement, err := txs.Prepare(query);
	if err != nil {
		txs.Rollback()
		return err
	}

	defer statement.Close()

	for _, item := range items {
		if item.Quantity <= 0 {
			txs.Rollback();
			return errors.New("quantity can not be equals to zero")
		}

		_, err := statement.Exec(
			inventoryId, 
			item.ProductVariantID,
			item.Quantity,
		)
		if err != nil {
			txs.Rollback()
			return err;
		}

	}
		return txs.Commit()
}