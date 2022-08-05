package core

import (
	"context"

	"github.com/gin-gonic/gin"
)

func RequestContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rc := NewContext(c)
		ctx := context.WithValue(c.Request.Context(), RequestContextKey, rc)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ForContext(ctx context.Context) *RequestContext {
	raw, _ := ctx.Value(RequestContextKey).(*RequestContext)
	return raw
}
