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