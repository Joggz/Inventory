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

	routes.InventoryRoutes(fb, h)


	return &App{fb: fb, db: db }
}

func (a *App) Listen(addr string) error {
	return a.fb.Listen(addr)
}


