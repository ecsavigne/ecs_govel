package grpcserverinit

import (
	conn "ecs_govel/internal/grpcservice/gen/services/v1/servicesv1connect"
	"ecs_govel/pkg/pkggin"
	"ecs_govel/pkg/pkglog"
	"ecs_govel/pkg/pkgmetrics/sdkopentelemetry"
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
	defer func() {
		if r := recover(); r != nil {
			logMessage := "Error in;.........  " + fmt.Sprintf("%+v", r)
			// panic(logMessage)
			pkglog.Log.Errorf("[ InitGrpcService ] - %s", logMessage)
		}
	}()

	if c_.GRPC_SERVER_PORT == "" {
		panic("not is possible start server grpc, server port is empty")
	}

	interceptors := connect.WithInterceptors(
		validate.NewInterceptor(),
		sdkopentelemetry.GetOtelInterceptor(),
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

	routerGin := pkggin.GetEngine()

	globalsMiddleware(routerGin)
	// routes ServiceHandler connectrpc
	routerGin.Any(path+"/*any", gin.WrapH(handler))
	// routes for anotations proto
	routerGin.Any("/productsapi/*any", gin.WrapH(transcoder))

	// addr := fmt.Sprintf("localhost:%s", c_.GRPC_SERVER_PORT) // comunicacion cerrada entre docker en la red de docker
	addr := fmt.Sprintf(":%s", c_.GRPC_SERVER_PORT) // external comunica con red de docker
	s := http.Server{
		Addr:      addr,
		Handler:   routerGin,
		Protocols: pkggin.CreateProtoHTTP2NotTLS(),
	}

	pkglog.Log.Sub("server").Infof("Start rpc server at: %s\n", addr)
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
