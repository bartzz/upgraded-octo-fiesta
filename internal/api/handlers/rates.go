package handlers

import (
	"kryptonim-interview/internal/api/middleware"
	"net/http"

	"kryptonim-interview/internal/service"

	"github.com/gin-gonic/gin"
)

type RatesHandler struct {
	svc *service.RatesService
}

func NewRatesHandler(svc *service.RatesService) *RatesHandler {
	return &RatesHandler{svc: svc}
}

func (h *RatesHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/rates",
		middleware.ValidateCurrenciesQueryParam(),
		h.HandleRates,
	)
}

func (h *RatesHandler) HandleRates(c *gin.Context) {
	currencies := middleware.GetCurrenciesFromContext(c)
	pairs, err := h.svc.GetAllPairs(c.Request.Context(), currencies)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, pairs)
}
