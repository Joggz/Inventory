package app

import (
	"database/sql"
	"vending-machine/internal/database"
	"vending-machine/internal/handlers"
	"vending-machine/internal/repository"
	"vending-machine/internal/routes"
	"vending-machine/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	fb *fiber.App
	db *sql.DB
	
}

func New() *App {
	fb := fiber.New(fiber.Config{
		AppName: "vending-machine",
	})
	fb.Use(logger.New())
	fb.Use(recover.New())

	db, err := database.ConnectDB();
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	repo := repository.NewInventoryRepository(db)
	svc := services.NewInventoryService(repo)
	h := handlers.NewInventoryHandler(svc)

	
	productRepo := repository.NewProductRepository(db)
	productVariantRepo := repository.NewProductVariantRepository(db)
	productSvc := services.NewProductService(productRepo, productVariantRepo)
	productHandler := handlers.NewProductHandler(productSvc)



	

	invns := repository.NewInventoryStockRepository(db);
	invsvc := services.NewInventoryStockService(invns);
	invstockHandler := handlers.NewInventoryStockHandler(invsvc)


	orderRepo := repository.NewOrderRepository(db)
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db, invns, orderRepo)
	orderHandler := handlers.NewOrderHandler(purchaseOrderRepo)
	

	routes.InventoryRoutes(fb, h)
	routes.ProductRoutes(fb, productHandler)
	routes.InventoryStockRoutes(fb, invstockHandler)
	routes.PurchaseStockRoutes(fb, orderHandler)


	return &App{fb: fb, db: db }
}

func (a *App) Listen(addr string) error {
	return a.fb.Listen(addr)
}


