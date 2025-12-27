package services

import (
	"vending-machine/internal/migrations"
	"vending-machine/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
	repoVariant *repository.ProductVariantRepository
	
}


func NewProductService(repo *repository.ProductRepository, repoVariant *repository.ProductVariantRepository) *ProductService {
	return &ProductService{
		repo: repo,
		repoVariant: repoVariant,
	}
}


func (p *ProductService) CreateProduct(product ProductInfo)(int64, error)  {
	productID, err :=p.repo.CreateProduct(migrations.CreateProduct{
		Name : product.Name,
		Description: product.Description,
	})

	if err != nil {
		return 0, err
	}
	var variants []migrations.CreateProductVariant

	for _, v := range product.Variants {
	 variants =  append(variants, migrations.CreateProductVariant{
		ProductID: productID,
		SKU: v.SKU,
		Name: v.Name,
		// Flavour: v.Flavour,
	 })
	}

	if err := p.repoVariant.CreateProductVariant(variants ); err != nil {
		return 0, err
	}
	return productID, nil
}
