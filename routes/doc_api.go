package routes

import (
	// "new_whatsmeow_gin/app/http/controllers"

	c_ "ecs_govel/configs"
	"ecs_govel/docs"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewProxy(protocol, targetHost string) (*httputil.ReverseProxy, error) {
	targetHost = fmt.Sprintf("%s://%s", protocol, targetHost)
	target, e := url.Parse(targetHost)
	if e != nil {
		return nil, errors.Errorf("Error parsing target host: %s Error: %s", targetHost, e.Error())
	}
	return httputil.NewSingleHostReverseProxy(target), nil
}

func RequestProxy(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func loadDocRoutes(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""

	proxy, e := NewProxy("http", fmt.Sprintf("%s:%s", c_.HTTP_SERVER_HOST, c_.HTTP_SERVER_PORT))
	if e != nil {
		log.Println(e)
		return
	}

	// Servir rutas swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	g.GET("/test1", RequestProxy(proxy))

	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
