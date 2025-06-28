package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
