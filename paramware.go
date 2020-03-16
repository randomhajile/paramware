package paramware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Int64Param adds an int param to the gin context.
func Int64Param(paramName string, defaultValue int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		if len(queryParams[paramName]) == 0 {
			c.Set(paramName, defaultValue)
			return
		}

		value, err := strconv.ParseInt(queryParams[paramName][0], 10, 64)
		if err != nil {
			err := fmt.Errorf("expected int %s param", paramName)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Set(paramName, value)
	}
}

// StringParam adds an int param to the gin context.
func StringParam(paramName string, defaultValue string) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		if len(queryParams[paramName]) == 0 {
			c.Set(paramName, defaultValue)
			return
		}

		c.Set(paramName, queryParams[paramName][0])
	}
}
