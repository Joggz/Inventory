package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"vending-machine/internal/handlers"
	"vending-machine/internal/repository"
	"vending-machine/internal/routes"
	"vending-machine/internal/services"
)



//   one particular item could have a variety
//  for instance yoghurt ( sweet, sour , other flavours, )
//  each of this should be identified  with an id specific to them,  
//  we should be able to restock a particular one, as buyer can buy out a particular flavour of yoghurt


// Task
//  1. create inventory , struct type
//  2. stock inventory, consider the above scenerio, update inventory table 
//  3. purchase item on inventory, deduct item count on table  when item 0, request restock.
// 4. restock inventory.  (another case admin designate task to be done, sends notification).
type App struct {
	fb *fiber.App
	
}

func New() *App {
	fb := fiber.New(fiber.Config{
		AppName: "vending-machine",
	})
	fb.Use(logger.New())
	fb.Use(recover.New())

	repo := repository.NewMemoryRepository()
	seedInventory(repo)
	svc := services.NewVendingService(repo)
	h := handlers.NewVendingHandler(svc)
	routes.Register(fb, h)

	return &App{fb: fb}
}

func (a *App) Listen(addr string) error {
	return a.fb.Listen(addr)
}

func seedInventory(r repository.Repository) {
	_ = r.SaveItem(&repository.Item{
		ID:         "cola",
		Name:       "Cola",
		PriceCents: 150,
		Quantity:   10,
	})
	_ = r.SaveItem(&repository.Item{
		ID:         "chips",
		Name:       "Chips",
		PriceCents: 200,
		Quantity:   7,
	})
	_ = r.SaveItem(&repository.Item{
		ID:         "water",
		Name:       "Water",
		PriceCents: 100,
		Quantity:   15,
	})
}

