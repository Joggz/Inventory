package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	// "errors"
	"vending-machine/internal/migrations"
	"vending-machine/internal/utils"
)

type PurchaseOrderRepository struct {
	db       *sql.DB
	invRepo  *InventoryStockRepository
	OrderRep *OrderRepository
}

func NewPurchaseOrderRepository(db *sql.DB, invRepo *InventoryStockRepository, orderRepo *OrderRepository) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{
		db:       db,
		invRepo:  invRepo,
		OrderRep: orderRepo,
	}
}

func (por *PurchaseOrderRepository) CreatePurchaseOrder(inventoryID, variantID, qty int64, email string) (*utils.PaystackInitializeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := por.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer tx.Rollback()

	qnty, price, err := por.invRepo.LockStock(tx, inventoryID, variantID)
	if err != nil {
		return nil, err
	}

	if qty > qnty {
		return nil, errors.New("insufficient stock")
	}
	amount := price * float64(qty) * 100
	if amount <= 0 {
		tx.Rollback()
		return nil, errors.New("we are not on promo abeg!!!")
	}
	fmt.Printf("amount should be: %f\n", amount)

	resp, err := utils.InitializePayment(email, int64(amount))
	if err != nil {
		tx.Rollback()
		// probably send emaiZl or return card
		return nil, err
	}

	order := migrations.Orders{
		InvoiceRef:       resp.Data.Reference,
		InventoryID:      inventoryID,
		ProductVariantID: variantID,
		Quantity:         int(qty),
	}

	if err := por.OrderRep.CreateOrder(tx, order); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return resp, nil

}


func (por *PurchaseOrderRepository) ConfirmPayment(ref string) ( *utils.PaymentConfirmation, error  ){
	res, err := utils.VerifyPaystack(ref);
	fmt.Printf("res: %v\n", res)
	if err != nil {
		return  nil, err
	}

	if !res.Data.Status {
		por.OrderRep.UpdateStatus(ref, "failed")
        return nil, errors.New("payment failed")
	}

	tx, err := por.db.Begin();
	if err != nil {
		return nil, err
	}

	defer tx.Rollback();

	var or struct {
		InventoryID int64
        VariantID   int64
        Quantity    int64
        Email       string
	}
	err = tx.QueryRow(`
        SELECT inventory_id, product_variant_id, quantity, email
        FROM orders WHERE reference = ?
    `, ref).Scan(&or.InventoryID, &or.VariantID, &or.Quantity, &or.Email)

    if err != nil {
        return nil, err
    }
	 err = por.invRepo.DeductStockUponSuccessfulPurchase(tx, or.InventoryID, or.VariantID, or.Quantity)
    if err != nil {
        return nil, err
    }

    por.OrderRep.UpdateStatus(ref, "success")
    tx.Commit()
    return res, nil
}