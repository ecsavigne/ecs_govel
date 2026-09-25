package restroute

import (
	"github.com/gin-gonic/gin"

	c_ "ecs_govel/configs"
	docs "ecs_govel/docs"
	"ecs_govel/pkg/pkgproxy"

	"github.com/ecsavigne/proxy-reverse/proxy"
)

func loadDocRoutes(g *gin.Engine) {
	// proxy reverse with server web api for test in docs
	pkgproxy.ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
		Host: "localhost",
		Port: c_.GRPC_SERVER_PORT,
	})
	g.Any("/ig/*any", pkgproxy.ReverseProxyApi.RequestProxy())

	// docs types
	g.GET(docs.DocHandler.DocsPath(), gin.WrapH(docs.DocHandler.DocsFunc()))

	// docs swagger
	g.GET(docs.DocHandler.SpecPath(), gin.WrapH(docs.DocHandler.SpecFunc()))

	if docs.DocHandler.AssetsEnabled() {
		g.GET(docs.DocHandler.AssetsPath()+"/*", func(c *gin.Context) {
			docs.DocHandler.Assets().ServeHTTP(c.Writer, c.Request)
		})
	}
}
