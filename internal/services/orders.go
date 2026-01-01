package services

import (
	// "database/sql"
	"database/sql"
	// "errors"
	"vending-machine/internal/migrations"
	"vending-machine/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (os *OrderService) CreateOrder(tx *sql.Tx, order migrations.Orders) error {
	return os.repo.CreateOrder(tx, order)
}
