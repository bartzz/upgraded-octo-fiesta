package middleware

import "github.com/gin-gonic/gin"

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
