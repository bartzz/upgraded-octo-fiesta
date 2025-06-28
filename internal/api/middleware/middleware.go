package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CtxCurrenciesKey = "currencies"
)

func GetCurrenciesFromContext(c *gin.Context) []string {
	if v, ok := c.Get(CtxCurrenciesKey); ok {
		if list, ok2 := v.([]string); ok2 {
			return list
		}
	}
	return nil
}

func ValidateCurrenciesQueryParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.Query("currencies")
		if raw == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		parts := strings.Split(raw, ",")
		if len(parts) < 2 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set(CtxCurrenciesKey, parts)
		c.Next()
	}
}
