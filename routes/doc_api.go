package routes

import (
	c_ "ecs_govel/configs"
	"ecs_govel/docs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadDocRoutes(g *gin.Engine) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", c_.ReverseProxyApi.Host, c_.ReverseProxyApi.Port)

	// Servir rutas swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	g.POST("/test", c_.ReverseProxyApi.RequestProxy())

	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
