package main

import (
	"kryptonim-interview/internal/api"
	"kryptonim-interview/internal/config"
	"kryptonim-interview/internal/net/rates/openexchangerates"
	"kryptonim-interview/internal/service"
	"log"
)

func main() {
	conf := config.Load()

	// Providers
	ratesProvider := openexchangerates.NewClient(conf.OpenExchangeRatesAPIKey)

	// Services
	ratesService := service.NewRatesService(ratesProvider)
	exchangeService := service.NewExchangeService()

	// Router
	router := api.SetupRouter(ratesService, exchangeService)

	// Start HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
