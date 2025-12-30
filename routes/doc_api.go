package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	c_ "ecs_govel/configs"
	_ "ecs_govel/docs"

	"github.com/ecsavigne/proxy-reverse/proxy"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadDocRoutes(g *gin.Engine) {
	c_.ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
		Host: "localhost",
		Port: "8080",
	})

	// Servir rutas swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	g.Any("/productsapi/*any", c_.ReverseProxyApi.RequestProxy())
	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
