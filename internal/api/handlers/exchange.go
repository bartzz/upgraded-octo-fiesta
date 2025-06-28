package handlers

import (
	"net/http"
	"strconv"

	"kryptonim-interview/internal/service"

	"github.com/gin-gonic/gin"
)

type ExchangeHandler struct {
	svc *service.ExchangeService
}

func NewExchangeHandler(svc *service.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{svc: svc}
}

func (h *ExchangeHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/exchange",
		h.HandleExchange,
	)
}

func (h *ExchangeHandler) HandleExchange(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result, err := h.svc.Exchange(from, to, amount)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"from": from, "to": to, "amount": result})
}
