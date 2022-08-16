package paramware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetParam generic function that can be used to get any kind of parameter.
func SetParam[T any](paramName string, defaultValue T, parser func(string) (T, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		if len(queryParams[paramName]) == 0 {
			c.Set(paramName, defaultValue)
			return
		}

		value, err := parser(queryParams[paramName][0])
		if err != nil {
			err := fmt.Errorf("expected %T %s param", *new(T), paramName)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Set(paramName, value)
	}
}

// Int64Param adds an int param to the gin context.
func Int64Param(paramName string, defaultValue int64) gin.HandlerFunc {
	parser := func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	}

	return SetParam(paramName, defaultValue, parser)
}

// StringParam adds a string param to the gin context.
func StringParam(paramName string, defaultValue string) gin.HandlerFunc {
	parser := func(s string) (string, error) {
		return s, nil
	}

	return SetParam(paramName, defaultValue, parser)
}

// BoolParam adds a boolean param to the gin context.
func BoolParam(paramName string, defaultValue bool) gin.HandlerFunc {
	return SetParam(paramName, defaultValue, strconv.ParseBool)
}
