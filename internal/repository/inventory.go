package repository

import (
	// "context"
	"database/sql"
	// "errors"
	// "fmt"
	// "time"

	"vending-machine/internal/migrations"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository  {
	return &InventoryRepository{db: db}
}

type InventoryService struct {
	Inventory func(db *sql.DB) ([]*migrations.Inventory, error)
	GetInventoryByID func(db *sql.DB, id int64)(*migrations.Inventory, error)
	UnInstallInventory func(db *sql.DB, id int64) error
	InstallInventory func(db *sql.DB, inventory []migrations.Inventory) error
}


func (r *InventoryRepository)Inventory()([]migrations.Inventory, error) {
	rows, err := r.db.Query("SELECT * FROM inventory")
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

func (r *InventoryRepository)GetInventoryByID(id int64)(*migrations.Inventory, error) {
	row := r.db.QueryRow("SELECT * FROM inventory WHERE id = ?", id);
	var i migrations.Inventory;
	if err := row.Scan(&i.ID, &i.Name, &i.Location, &i.CreatedAt); err != nil {
		return nil, err;
	}
	return &i, nil
}

func (r *InventoryRepository)UnInstallInventory(id int64) error {
	_, err := r.db.Exec("DELETE FROM inventory WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}


func (r *InventoryRepository) InstallInventory(inventory []migrations.Inventory) error {
	if len(inventory) == 0 {
		return  nil
	};

    inventoryQuery := `
    INSERT INTO inventories (name, location)
    VALUES (?, ?);
    `

	tx, err := r.db.Begin();
	if err != nil {
		return err
	}

	statement, err := tx.Prepare(inventoryQuery);
	if err != nil {
		tx.Rollback();
		return err
	}

	defer statement.Close();


    for _, i := range inventory {
        _, err := statement.Exec( i.Name, i.Location)	
        if err != nil {
			tx.Rollback()
            return err
        }
    }
    return tx.Commit()
}