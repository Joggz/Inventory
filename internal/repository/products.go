package repository;

import (
	"database/sql"
	"vending-machine/internal/migrations"
)


type ProductService struct {
	Products func(db *sql.DB) ([]*migrations.Product, error)
	GetProductByID func(db *sql.DB, id int64)(*migrations.Product, error)
	RemoveProductByID func(db *sql.DB, id int64) error
}



func Products(db *sql.DB)([]migrations.Product, error)  {
	rows, err := db.Query("SELECT * From product");

	if err != nil {
		return nil, err
	}

	var products []migrations.Product

	for rows.Next(){
		var p migrations.Product;
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	return  products, nil
}

func GetProductByID(db *sql.DB, id int64)(*migrations.Product, error)  {
	row := db.QueryRow(`SELECT * FROM product WHERE id = ?`, id);
	var p migrations.Product;
	if err:= row.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt);  err != nil {
		return nil, err
	}
	return &p, nil
}

func RemoveProductByID(db *sql.DB, id int64) error  {
	_, err := db.Exec(`DELETE FROM product WHERE id = ?`, id);
	if err != nil {
		return  err
	}
	return nil;
}