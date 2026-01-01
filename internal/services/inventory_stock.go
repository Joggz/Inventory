package services

import (
	// "database/sql"
	"errors"
	"vending-machine/internal/migrations"
	"vending-machine/internal/repository"
)

type InventoryStockService struct {
	repo *repository.InventoryStockRepository
}

func NewInventoryStockService(repo *repository.InventoryStockRepository) *InventoryStockService {
	return &InventoryStockService{repo: repo}
}


func (invs *InventoryStockService) AddMultipleStock( inventoryId int64, items [] migrations.AddStockItem) error {
	if len(items) == 0 {
		return errors.New("no stock items provided")
	}
	if inventoryId == 0 {
		return errors.New("invalid inventory")
	}

	return invs.repo.AddMmultiStock(inventoryId, items)
}

func (invs *InventoryStockService) GetAllStocks()([]migrations.Stocks, error)  {
		return invs.repo.GetAllStocks();
}

func (invs *InventoryStockService)GetStocksByInventoryID(inventoryID int64)([]migrations.Stocks, error)  {
	if inventoryID <= 0 {
		return nil, errors.New("invalid inventory")
	}
	
	return invs.repo.GetStocksByInventoryID(inventoryID);
}

func (invs *InventoryStockService) GetProductVariantByInventoryID(inventoryID int64, productVariant int64)([]migrations.Stocks, error)  {
	if inventoryID <= 0 {
		return nil, errors.New("invalid inventory")
	}
	if productVariant <= 0 {
		return nil, errors.New("invalid product variant")
	}
	return invs.repo.GetProductVariantByInventoryID(inventoryID, productVariant);
}