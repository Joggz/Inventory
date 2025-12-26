package services

import (
	"vending-machine/internal/migrations"
	"vending-machine/internal/repository"
	"errors"
)

type InventoryService struct {
	repo *repository.InventoryRepository
}

func NewInventoryService(repo *repository.InventoryRepository) *InventoryService{
	return &InventoryService{repo: repo}
}

func(s *InventoryService) Create(inventory []migrations.Inventory) error {
	if len(inventory) == 0 {
		return  errors.New("inventory is empty")
	}

	err := s.repo.InstallInventory(inventory)
	if err != nil {
		return err
	}
	return nil
}


func (s *InventoryService) GetAllInventory() ([]migrations.Inventory, error) {
	invs, err := s.repo.Inventory()
	if err != nil {
		return nil, err
	}
	return invs, nil
}