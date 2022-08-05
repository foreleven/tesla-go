package core

import (
	"errors"
	"strings"
	"tesla-go/tesla"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

type RequestContext struct {
	token      *AuthToken
	client     *tesla.Client
	ginContext *gin.Context
}

var RequestContextKey = &contextKey{"tesla-go-request-context"}

func NewContext(c *gin.Context) *RequestContext {
	authorization := c.Request.Header.Get("Authorization")
	context := &RequestContext{
		ginContext: c,
	}

	if authorization == "" {
		return context
	} else {
		parts := strings.SplitN(authorization, " ", 2)
		if parts[0] != "Bearer" {
			return &RequestContext{
				ginContext: c,
			}
		}
		tokenString := parts[1]
		token, err := Verify(tokenString)
		if err != nil {
			return context
		}
		context.token = token
		client, _ := tesla.NewClient(tesla.ClientOptions{
			AuthToken: tesla.AuthToken{AccessToken: token.payload.AccessToken},
		})
		context.client = client
	}

	return context
}

func (c RequestContext) HasAuthorized() bool {
	if c.token == nil {
		c.ginContext.AbortWithError(403, errors.New("Forbidden"))
		return false
	}
	return true
}

func (c RequestContext) GetClient() *tesla.Client {
	return c.client
}
