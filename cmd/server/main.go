package main

import (
	"log"
	"os"

	"vending-machine/internal/app"
	"vending-machine/internal/database"
)

func main() {
	a := app.New()
	db, err := database.ConnectDB()
	if err != nil {
		log.Printf("database unavailable: %v (running with in-memory repository)", err)
	} else {
		defer db.Close()
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("vending-machine listening on :%s\n", port)
	if err := a.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
