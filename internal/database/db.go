package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"vending-machine/internal/migrations"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	// pass := os.Getenv("DB_PASSWORD")
	pass := "Qwertyuiop0!";
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "inventory_app_db"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Error connecting to database: %w", err)
	}

	if err := migrations.CreateInventoryTable(db); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Error creating inventory table: %w", err)
	}

	if err := migrations.CreateProductTable(db); err != nil {
		_ = db.Close();
		return nil, fmt.Errorf("Error creating product table: %w", err)
	}

	if err := migrations.CreateProductVariantTable(db); err != nil {
		_ = db.Close();
		return nil, fmt.Errorf("Error creating product variant table: %w", err)
	}
	if err := migrations.CreateInventoryStockTable(db); err != nil {
		_ = db.Close();
		return nil, fmt.Errorf("Error creating inventory stock table: %w", err)
	}
	
	fmt.Println("✅ MySQL connected successfully")
	return db, nil
}
