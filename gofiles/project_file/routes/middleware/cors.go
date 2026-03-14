package middleware

import (
	"fmt"

	c_ "ecs_govel/configs"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// origins allowed
		domains := []struct {
			Host     string
			Port     string
			Protocol string
		}{
			{c_.HTTP_SERVER_HOST_DOC_API, c_.HTTP_SERVER_PORT_DOC_API, "http"},
		}

		for _, domain := range domains {
			c.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%s://%s:%s", domain.Protocol, domain.Host, domain.Port))
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
