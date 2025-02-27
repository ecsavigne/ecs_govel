package routes

import (
	c_ "ecs_govel/configs"
	"ecs_govel/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadDocRoutes(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""

	// Servir rutas swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	g.POST("/test", c_.ReverseProxyApiDoc.RequestProxy())

	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
