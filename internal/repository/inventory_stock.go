package repository

import (
	"database/sql"
	"errors"
	"vending-machine/internal/migrations"
)

type InventoryStockRepository struct {
	db *sql.DB
}

type AddStockItem struct {
	ProductVariantID int64
	quantity         int
}

func NewInventoryStockRepository(db *sql.DB) *InventoryStockRepository {
	return &InventoryStockRepository{db: db}
}

func (invs *InventoryStockRepository) AddMmultiStock(inventoryId int64, items []migrations.AddStockItem) error {
	if len(items) == 0 {
		return nil
	}

	query := `
	INSERT INTO inventory_stock (inventory_id, product_variant_id, quantity, price)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		quantity = quantity + VALUES(quantity),
		price = VALUES(price)
	`

	txs, err := invs.db.Begin()
	if err != nil {
		txs.Rollback()
		return err
	}

	statement, err := txs.Prepare(query)
	if err != nil {
		txs.Rollback()
		return err
	}

	defer statement.Close()

	for _, item := range items {
		println(item.Price)
		if item.Quantity <= 0 {
			txs.Rollback()
			return errors.New("quantity can not be equals to zero")
		}
		if item.Price == 0 {
			txs.Rollback()
			return errors.New("price can not be equals to zero")
		}
		_, err := statement.Exec(
			inventoryId,
			item.ProductVariantID,
			item.Quantity,
			item.Price,
		)
		if err != nil {
			txs.Rollback()
			return err
		}

	}
	return txs.Commit()
}

func (invs *InventoryStockRepository) GetAllStocks() ([]migrations.Stocks, error) {
	query := `
		SELECT
		i.id,
		i.name,
		p.id,
		p.name,
		pv.id,
		pv.flavour,
		pv.sku,
		s.quantity
	FROM inventory_stock s
	JOIN inventories i ON i.id = s.inventory_id
	JOIN product_variants pv ON pv.id = s.product_variant_id
	JOIN products p ON p.id = pv.product_id
	ORDER BY i.id, p.id, pv.id
	`

	rows, err := invs.db.Query(query)
	if err != nil {
		return nil, err
	}
	var stock []migrations.Stocks

	for rows.Next() {
		var s migrations.Stocks
		err := rows.Scan(
			&s.InventoryID,
			&s.InventoryName,
			&s.ProductID,
			&s.ProductName,
			&s.VariantID,
			&s.VariantName,
			&s.SKU,
			&s.Quantity,
			&s.Price,
		)

		if err != nil {
			return nil, err
		}
		stock = append(stock, s)
	}
	return stock, nil
}

func (invs *InventoryStockRepository) GetStocksByInventoryID(inventoryID int64) ([]migrations.Stocks, error) {
	query := `
	SELECT
		i.id,
		i.name,
		p.id,
		p.name,
		pv.id,
		pv.name,
		pv.sku,
		s.quantity
	FROM inventory_stock s
	JOIN inventories i ON i.id = s.inventory_id
	JOIN product_variants pv ON pv.id = s.product_variant_id
	JOIN products p ON p.id = pv.product_id
	WHERE s.inventory_id = ?
	ORDER BY p.id, pv.id
	`

	rows, err := invs.db.Query(query, inventoryID)
	if err != nil {
		return nil, err
	}
	var stock []migrations.Stocks

	for rows.Next() {
		var s migrations.Stocks
		err := rows.Scan(
			&s.InventoryID,
			&s.InventoryName,
			&s.ProductID,
			&s.ProductName,
			&s.VariantID,
			&s.VariantName,
			&s.SKU,
			&s.Quantity,
		)

		if err != nil {
			return nil, err
		}
		stock = append(stock, s)
	}
	return stock, nil
}

func (invs *InventoryStockRepository) GetProductVariantByInventoryID(inventoryID int64, productVariant int64) ([]migrations.Stocks, error) {
	query := `
	SELECT
		i.id AS inventory_id,
		i.name AS inventory_name,
		p.id AS product_id,
		p.name AS product_name,
		pv.id AS variant_id,
		pv.sku AS variant_sku,
		s.quantity,
		s.price
	FROM inventory_stock s
	JOIN inventories i ON i.id = s.inventory_id
	JOIN product_variants pv ON pv.id = s.product_variant_id
	JOIN products p ON p.id = pv.product_id
	WHERE s.inventory_id = ?
	AND pv.id = ?
	ORDER BY p.id, pv.id
	`

	rows, err := invs.db.Query(query, inventoryID, productVariant)
	if err != nil {
		return nil, err
	}
	var stock []migrations.Stocks

	for rows.Next() {
		var s migrations.Stocks
		err := rows.Scan(
			&s.InventoryID,
			&s.InventoryName,
			&s.ProductID,
			&s.ProductName,
			&s.VariantID,
			// &s.VariantName,
			&s.SKU,
			&s.Quantity,
			&s.Price,
		)

		if err != nil {
			return nil, err
		}
		stock = append(stock, s)
	}
	return stock, nil
}

func (invs *InventoryStockRepository) LockStock(tx *sql.Tx, inventory_id int64, variant_id int64) (quantity int64, price float64, err error) {
	err = tx.QueryRow(`
        SELECT quantity, price
        FROM inventory_stock
        WHERE inventory_id = ? AND product_variant_id = ?
        FOR UPDATE
    `, inventory_id, variant_id).Scan(&quantity, &price)

	if err != nil {
		return 0, 0, err
	}

	return quantity, price, nil
}

func (invs *InventoryStockRepository) DeductStockUponSuccessfulPurchase(tx *sql.Tx, inventory_id int64, quantity int64, variant_id int64) error {
	_, err := tx.Exec(`
        UPDATE inventory_stock
        SET quantity = quantity - ?
        WHERE inventory_id = ? AND product_variant_id = ?
    `, quantity, inventory_id, variant_id)

	return err
}
