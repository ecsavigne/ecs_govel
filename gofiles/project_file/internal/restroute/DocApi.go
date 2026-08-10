package restroute

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	c_ "ecs_govel/configs"
	docs "ecs_govel/docs"
	"ecs_govel/pkg/pkgproxy"

	"github.com/ecsavigne/proxy-reverse/proxy"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadDocRoutes(g *gin.Engine) {
	// proxy reverse with server web api for test in docs
	pkgproxy.ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
		Host: "localhost",
		Port: c_.GRPC_SERVER_PORT,
	})
	g.Any("/name_api/*any", pkgproxy.ReverseProxyApi.RequestProxy())

	// docs types
	switch strings.ToLower(c_.TYPE_DOCUMENTATION) {
	case "swagger":

		// Servir rutas swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		g.GET("/docs", func(gc *gin.Context) {
			gc.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
	case "scalar":
		g.GET("/docs", func(gc *gin.Context) {
			//c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			gc.Data(http.StatusOK, "text/html; charset=utf-8", []byte(docs.ScalarHtml))
		})
	}
}
