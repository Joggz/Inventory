package repository

import (
	"database/sql"
	"vending-machine/internal/migrations"
)

type ProductVariantRepository struct {
	db *sql.DB
}

func NewProductVariantRepository(db *sql.DB) *ProductVariantRepository  {
	return &ProductVariantRepository{db: db}
}


func (pv *ProductVariantRepository) CreateProductVariant(pdv []migrations.CreateProductVariant) error {
	if len(pdv) == 0 {
		return nil
	};

	query := `INSERT INTO product_variants (product_id, flavour, sku) VALUES (?, ?, ?)`;

	tx, err := pv.db.Begin()
	if err != nil {
		return err
	}

	statement, err := tx.Prepare(query);
	if err != nil {
		tx.Rollback()
		return err
	}
	defer statement.Close();

	for _, v := range pdv {
		_, err := statement.Exec(
			v.ProductID,
			v.Flavour,
			v.SKU,
		)
		if err != nil {
			tx.Rollback();
			return err
		}
	}
	return tx.Commit();
}