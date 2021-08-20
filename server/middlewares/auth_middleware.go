package middlewares

import (
	"github.com/IgorLomba/API-REST-GO/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			// unatorhized
			c.AbortWithStatus(401)
		}
		// get the size of Bearer_schema (Bearer [GET THIS PART])
		token := header[len(Bearer_schema):]

		// if token is not valid
		if !services.NewJwtService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
