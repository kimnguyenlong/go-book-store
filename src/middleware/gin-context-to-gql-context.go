package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinContextToGQLContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContext", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
