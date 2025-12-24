package services

import (
	// "context"
	"database/sql"
	// "errors"
	// "fmt"
	// "time"

	"vending-machine/internal/migrations"
)


type InventoryService struct {
	Inventory func(db *sql.DB) ([]*migrations.Inventory, error)
	GetInventoryByID func(db *sql.DB, id int64)(*migrations.Inventory, error)
	UnInstallInventory func(db *sql.DB, id int64) error
}


func Inventory(db *sql.DB)([]migrations.Inventory, error) {
	rows, err := db.Query("SELECT * FROM inventory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var inventories []migrations.Inventory;

	for rows.Next() {
		var i migrations.Inventory;
		if err := rows.Scan(&i.ID, &i.Name, &i.Location, &i.CreatedAt); err != nil {
			return nil, err;
		}
		inventories = append(inventories, i);

	}
	return inventories, nil
}

func GetInventoryByID(db *sql.DB, id int64)(*migrations.Inventory, error) {
	row := db.QueryRow("SELECT * FROM inventory WHERE id = ?", id);
	var i migrations.Inventory;
	if err := row.Scan(&i.ID, &i.Name, &i.Location, &i.CreatedAt); err != nil {
		return nil, err;
	}
	return &i, nil
}

func UnInstallInventory(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM inventory WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}