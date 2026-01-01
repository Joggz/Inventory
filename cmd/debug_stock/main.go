package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	// Fallback to default if envs are empty (just in case, matching db.go logic partially)
	if user == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/inventory_app_db?parseTime=true"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")

	var quantity int
	var price float64
	inventoryID := 1
	variantID := 2

	query := `
		SELECT quantity, price
		FROM inventory_stock
		WHERE inventory_id = ? AND product_variant_id = ?
	`
	fmt.Printf("Querying for inventory_id=%d, variant_id=%d\n", inventoryID, variantID)
	
	err = db.QueryRow(query, inventoryID, variantID).Scan(&quantity, &price)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
	} else {
		fmt.Printf("Found: Quantity=%d, Price=%f\n", quantity, price)
	}

	// List all stocks to be sure
	rows, err := db.Query("SELECT inventory_id, product_variant_id, quantity, price FROM inventory_stock")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("All stocks:")
	for rows.Next() {
		var iID, vID int
		var q int
		var p float64
		if err := rows.Scan(&iID, &vID, &q, &p); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("InvID: %d, VarID: %d, Qty: %d, Price: %f\n", iID, vID, q, p)
	}
}
