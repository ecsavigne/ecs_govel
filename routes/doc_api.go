package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	c_ "ecs_govel/configs"
	docs "ecs_govel/docs"

	"github.com/ecsavigne/proxy-reverse/proxy"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadDocRoutes(g *gin.Engine) {
	// proxy reverse with server web api for test in docs
	c_.ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
		Host: "localhost",
		Port: c_.GRPC_SERVER_PORT,
	})
	g.Any("/productsapi/*any", c_.ReverseProxyApi.RequestProxy())

	// docs types
	switch strings.ToLower(c_.TYPE_DOCUMENTATION) {
	case "swagger":

		// Servir rutas swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		g.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
	case "scalar":
		g.GET("/docs", func(c *gin.Context) {
			//c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(docs.ScalarHtml))
		})
	}
}
