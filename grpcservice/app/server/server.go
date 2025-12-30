package grpcserverinit

import (
	conn "ecs_govel/grpcservice/gen/services/v1/servicesv1connect"
	"fmt"
	"net/http"
	"strings"

	c_ "ecs_govel/configs"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/Cyprinus12138/otelgin"
	"github.com/gin-gonic/gin"
)

func globalsMiddleware(g *gin.Engine) {
	// register instrumentation of metrics and tracing
	g.Use(otelgin.Middleware("productsapi"))

	g.Use(func(c *gin.Context) {
		fmt.Printf("Call request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
}

func InitGrpcService() {
	if c_.GRPC_SERVER_PORT == "" {
		return
	}

	interceptors := connect.WithInterceptors(
		validate.NewInterceptor(),
		c_.GetOtelInterceptor(),
	)

	path, handler := conn.NewProductServiceHandler(
		new(ProductService),
		// Validation via Protovalidate is almost always recommended
		interceptors,
	)

	// Traqnscodificador
	serviceVanguard := vanguard.NewService(path, handler)
	transcoder, err := vanguard.NewTranscoder([]*vanguard.Service{
		serviceVanguard,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create transcoder: %v", err))
	}

	if strings.ToLower(c_.APP_MODE) == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	routerGin := c_.GetEngine()

	globalsMiddleware(routerGin)

	// routes ServiceHandler connectrpc
	routerGin.Any(path+"/*any", gin.WrapH(handler))
	// routes for anotations proto
	routerGin.Any("/productsapi/*any", gin.WrapH(transcoder))

	addr := fmt.Sprintf("localhost:%s", c_.GRPC_SERVER_PORT)
	s := http.Server{
		Addr:      addr,
		Handler:   routerGin,
		Protocols: c_.CreateProtoHTTP2NotTLS(),
	}

	c_.Log.Sub("server").Infof("Start rpc server at: %s\n", addr)
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
