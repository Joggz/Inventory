package migrations

import "database/sql";

func CreateInventoryTable(db *sql.DB) error  {
	query := `CREATE TABLE IF NOT EXISTS inventories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_inventory_name (name)
);
`
	_, err := db.Exec(query)
	return err
}

func CreateProductTable(db *sql.DB) error {
	productQuery := `
	CREATE TABLE IF NOT EXISTS products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
	_, err := db.Exec(productQuery)
	return err
}

func CreateProductVariantTable(db *sql.DB) error {
	variantQuery := `
	CREATE TABLE IF NOT EXISTS product_variants (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT NOT NULL,
    flavour VARCHAR(50) NOT NULL,
    sku VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
`
	_, err := db.Exec(variantQuery)
	return err
}


func CreateInventoryStockTable(db *sql.DB) error {
	stockQuery := `
	CREATE TABLE IF NOT EXISTS inventory_stock (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    inventory_id BIGINT NOT NULL,
    product_variant_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY uniq_inventory_variant (inventory_id, product_variant_id),

    FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE,
    FOREIGN KEY (product_variant_id) REFERENCES product_variants(id) ON DELETE CASCADE
);
`
	_, err := db.Exec(stockQuery);
	return err;
}


// func AddPriceToInventoryStock(db *sql.DB) error  {
//     query := `
//     ALTER TABLE inventory_stock
//     ADD COLUMN price DECIMAL(10, 2) NOT NULL DEFAULT 0.00;
//     `
//     _, err := db.Exec(query)
//     return err
// }

func AddPriceToInventoryStock(db *sql.DB) error {
	// Check if column already exists
	checkQuery := `
		SELECT COUNT(*)
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA = DATABASE()
		AND TABLE_NAME = 'inventory_stock'
		AND COLUMN_NAME = 'price'
	`

	var count int
	if err := db.QueryRow(checkQuery).Scan(&count); err != nil {
		return err
	}

	// Add column only if it doesn't exist
	if count == 0 {
		alterQuery := `
			ALTER TABLE inventory_stock
			ADD COLUMN price DECIMAL(10,2) NOT NULL DEFAULT 0.00
		`
		_, err := db.Exec(alterQuery)
		return err
	}

	return nil
}

func CreatwOrderTable(db *sql.DB) error  {
	orderQuery := `
		CREATE TABLE IF NOT EXISTS orders (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		reference VARCHAR(100) UNIQUE,
		inventory_id BIGINT,
		product_variant_id BIGINT,
		quantity INT,
		amount BIGINT,
		email VARCHAR(255),
		status ENUM('pending','success','failed') DEFAULT 'pending',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_,  err := db.Exec(orderQuery);
	
	return err
}


// create table
// lock row stock on inventory stock
// create Order
// initiate paymnent
// verify payment
// deduct stock quantuity

