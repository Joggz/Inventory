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

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository  {
	return &ProductRepository{db: db}
}


func (p *ProductRepository) CreateProduct(data migrations.CreateProduct) (int64, error)  {
	
	query := `INSERT INTO products (name, description) VALUES (?, ?)`;
	result, err := p.db.Exec(query, data.Name, data.Description);
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}




func (p *ProductRepository)Products()([]*migrations.Product, error)  {
	rows, err := p.db.Query("SELECT * From products");

	if err != nil {
		return nil, err
	}

	var products []*migrations.Product

	for rows.Next(){
		var p migrations.Product;
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, err
		}

		products = append(products, &p)
	}
	return  products, nil
}

func (p *ProductRepository)GetProductByID(id int64)(*migrations.Product, error)  {
	row := p.db.QueryRow(`SELECT * FROM products WHERE id = ?`, id);
	var pdt migrations.Product;
	if err:= row.Scan(&pdt.ID, &pdt.Name, &pdt.Description, &pdt.CreatedAt);  err != nil {
		return nil, err
	}
	return &pdt, nil
}

func (p *ProductRepository)RemoveProductByID(id int64) error  {
	_, err := p.db.Exec(`DELETE FROM products WHERE id = ?`, id);
	if err != nil {
		return  err
	}
	return nil;
}