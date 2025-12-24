package repository

import (
	"errors"
	"sync"
)

var (
	ErrNotFound     = errors.New("item not found")
	ErrInvalidStock = errors.New("invalid stock")
)

type Item struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PriceCents int    `json:"price_cents"`
	Quantity   int    `json:"quantity"`
}

type Repository interface {
	GetAll() ([]*Item, error)
	GetByID(id string) (*Item, error)
	SaveItem(i *Item) error
	UpdateQuantity(id string, delta int) (*Item, error)
}

type memoryRepo struct {
	mu    sync.RWMutex
	items map[string]*Item
}

func NewMemoryRepository() Repository {
	return &memoryRepo{
		items: make(map[string]*Item),
	}
}

func (m *memoryRepo) GetAll() ([]*Item, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]*Item, 0, len(m.items))
	for _, it := range m.items {
		c := *it
		out = append(out, &c)
	}
	return out, nil
}

func (m *memoryRepo) GetByID(id string) (*Item, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	it, ok := m.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	c := *it
	return &c, nil
}

func (m *memoryRepo) SaveItem(i *Item) error {
	if i == nil || i.ID == "" {
		return errors.New("invalid item")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	c := *i
	m.items[i.ID] = &c
	return nil
}

func (m *memoryRepo) UpdateQuantity(id string, delta int) (*Item, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	it, ok := m.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	newQty := it.Quantity + delta
	if newQty < 0 {
		return nil, ErrInvalidStock
	}
	it.Quantity = newQty
	c := *it
	return &c, nil
}

