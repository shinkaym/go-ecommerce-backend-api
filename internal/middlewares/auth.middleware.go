package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shinkaym/go-ecommerce-backend-api/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "Invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
