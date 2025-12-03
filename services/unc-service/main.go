package main

import (
	"log"
	"unc/services/unc-service/setup"
)

func main() {
	app := setup.NewApp()

	if err := app.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
