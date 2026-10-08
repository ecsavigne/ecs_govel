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

	midleware "ecs_govel/internal/grpcservice/app/server/interceptor"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
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

	srvConnect := connect.NewServer(
		validate.NewServerInterceptor(),
		sdkopentelemetry.GetOtelInterceptor(),
	)

	conn.RegisterProductServiceHandler(srvConnect, new(ProductService))
	router := http.NewServeMux()
	connecthttp.Mount(router, srvConnect)

	path := fmt.Sprintf("/%s/", conn.ProductServiceName)

	// Traqnscodificador
	serviceVanguard := vanguard.NewService(path, router)
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
	// routes ServiceHandler connectrpc, grpc and grpc-WEB
	routerGin.Any(path+"/*any", midleware.CorsReq(), gin.WrapH(router))
	// routes for anotations proto

	// route rest with hash
	routerGin.Any("/"+c_.HASH_ROUTE+"/*any" /*interceptor.AdapterMiddleware(), */, gin.WrapH(transcoder))

	// addr := fmt.Sprintf("localhost:%s", c_.GRPC_SERVER_PORT) // comunicacion cerrada entre docker en la red de docker
	addr := fmt.Sprintf("%s:%s", c_.HTTP_SERVER_HOST, c_.GRPC_SERVER_PORT) // external comunica con red de docker
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
