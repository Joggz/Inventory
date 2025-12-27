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


type CreateProduct struct {
	Name        string
	Description string
}

type CreateProductVariant struct {
	ProductID int64
	Flavour   string
	SKU       string
	Name string
}


type CreateVariantInput struct {
	ProductID int64
	Name      string
	SKU       string
}


type AddStockItem struct  {
	ProductVariantID int64
	Quantity int
}

type AddInventoryStock  struct {
	InventoryID int64
	Items [] AddStockItem
}

type AddMultipleStockPayload struct {
	Items []struct {
		VariantID int64 `json:"variant_id"`
		Quantity  int   `json:"quantity"`
	} `json:"items"`
}
