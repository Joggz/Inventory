package migrations

import (
	"time"
)

type Inventory struct {
	ID        int64
	Name      string
	Location  string
	CreatedAt time.Time
}


type Product struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

type ProductVariant struct {
	ID        int64
	ProductID int64
	Flavour   string
	SKU       string
	CreatedAt time.Time
}

type InventoryStock struct {
	ID               int64
	InventoryID      int64
	ProductVariantID int64
	Quantity         int
	UpdatedAt        time.Time
}
