package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	// "github.com/ecsavigne/proxy-reverse/proxy"
	ginSwagger "github.com/swaggo/gin-swagger"
	// c_ "ecs_govel/configs"
	_ "ecs_govel/docs"
)

func loadDocRoutes(g *gin.Engine) {
	// c_.ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
	// 	Host: "localhost",
	// 	Port: "2233",
	// })

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s\n", c_.ReverseProxyApi.Host, c_.ReverseProxyApi.Port)
	// fmt.Println("docs.SwaggerInfo.Host : ", docs.SwaggerInfo.Host)

	// Servir rutas swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// g.POST("/test", c_.ReverseProxyApi.RequestProxy())

	// g.GET("/a", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Route Doc Root")
	// })

	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
