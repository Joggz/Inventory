package services

import (
	"errors"
	"vending-machine/internal/repository"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type PurchaseInput struct {
	ItemID    string
	Quantity  int
	PaidCents int
}

type PurchaseResult struct {
	TotalCents        int
	ChangeCents       int
	RemainingQuantity int
	Item              *repository.Item
}

type VendingService interface {
	Inventory() ([]*repository.Item, error)
	Purchase(in *PurchaseInput) (*PurchaseResult, error)
	Restock(itemID string, qty int) (*repository.Item, error)
}

type vendingService struct {
	repo repository.Repository
}

func NewVendingService(r repository.Repository) VendingService {
	return &vendingService{repo: r}
}

func (v *vendingService) Inventory() ([]*repository.Item, error) {
	return v.repo.GetAll()
}

func (v *vendingService) Purchase(in *PurchaseInput) (*PurchaseResult, error) {
	if in == nil || in.ItemID == "" || in.Quantity <= 0 {
		return nil, errors.New("invalid purchase input")
	}
	item, err := v.repo.GetByID(in.ItemID)
	if err != nil {
		return nil, err
	}
	total := item.PriceCents * in.Quantity
	if in.PaidCents < total {
		return nil, ErrInsufficientFunds
	}
	updated, err := v.repo.UpdateQuantity(in.ItemID, -in.Quantity)
	if err != nil {
		return nil, err
	}
	change := in.PaidCents - total
	return &PurchaseResult{
		TotalCents:        total,
		ChangeCents:       change,
		RemainingQuantity: updated.Quantity,
		Item:              item,
	}, nil
}

func (v *vendingService) Restock(itemID string, qty int) (*repository.Item, error) {
	if itemID == "" || qty <= 0 {
		return nil, errors.New("invalid restock input")
	}
	return v.repo.UpdateQuantity(itemID, qty)
}

