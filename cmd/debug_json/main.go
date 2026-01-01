package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Orders struct {
	ID            int64
	InvoiceRef    string
	InventoryID   int64
	ProductVariantID int64
	Quantity      int
	CreatedAt     time.Time
	Email         string
}

func main() {
	jsonStr := `{
   "invoice_ref": "INV-2025-0001", 
   "inventory_id": 1, 
   "product_variant_id": 2, 
   "quantity": 2, 
   "email": "customer@example.com" 
 }`

	var payload Orders
	err := json.Unmarshal([]byte(jsonStr), &payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed: InventoryID=%d, ProductVariantID=%d\n", payload.InventoryID, payload.ProductVariantID)
}
