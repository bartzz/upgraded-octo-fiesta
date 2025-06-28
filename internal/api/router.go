package api

import (
	"kryptonim-interview/internal/api/handlers"
	"kryptonim-interview/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRatesRouter(ratesSvc *service.RatesService) *gin.Engine {
	r := gin.Default()
	ratesH := handlers.NewRatesHandler(ratesSvc)
	ratesH.RegisterRoutes(r)
	return r
}

func SetupRouter(ratesSvc *service.RatesService, exchangeSvc *service.ExchangeService) *gin.Engine {
	r := gin.Default()

	// Rates handlers.
	ratesH := handlers.NewRatesHandler(ratesSvc)
	ratesH.RegisterRoutes(r)

	// Exchange handlers.
	exchangeH := handlers.NewExchangeHandler(exchangeSvc)
	exchangeH.RegisterRoutes(r)

	return r
}
