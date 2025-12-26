package grpcserverinit

import (
	"context"
	services "ecs_govel/grpcservice/gen/services/v1"
	conn "ecs_govel/grpcservice/gen/services/v1/servicesv1connect"
	"fmt"
	"net/http"
	"strings"

	c_ "ecs_govel/configs"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct{}

func (srv *ProductService) CreateProduct(ctx context.Context, reqWrapper *connect.Request[services.CreateProductRequest]) (*connect.Response[services.CreateProductResponse], error) {
	resp := &services.CreateProductResponse{}

	req := reqWrapper.Msg
	resp.SetProduct(req.GetProduct())
	msg := strings.Builder{}
	msg.WriteString("Product with name: " + req.GetProduct().GetName() + " has been created")

	resp.SetMsg(msg.String())

	return connect.NewResponse(resp), nil
}

func (srv *ProductService) GetProduct(ctx context.Context, reqWrapper *connect.Request[services.GetProductRequest]) (*connect.Response[services.GetProductResponse], error) {

	return connect.NewResponse(&services.GetProductResponse{}), nil
}

func (srv *ProductService) GetProducts(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[services.GetProductsResponse], error) {
	// get header
	fmt.Printf("X-User-ID: %s\n", req.Header().Get("X-User-ID"))

	return connect.NewResponse(&services.GetProductsResponse{}), nil
}

func (srv *ProductService) UpdateProduct(ctx context.Context, reqWrapper *connect.Request[services.UpdateProductRequest]) (*connect.Response[services.UpdateProductResponse], error) {

	return connect.NewResponse(&services.UpdateProductResponse{}), nil
}

func (srv *ProductService) DeleteProduct(ctx context.Context, reqWrapper *connect.Request[services.DeleteProductRequest]) (*connect.Response[services.DeleteProductResponse], error) {

	return connect.NewResponse(&services.DeleteProductResponse{}), nil
}

func globalsMiddleware(g *gin.Engine) {
	g.Use(func(c *gin.Context) {
		fmt.Printf("Call request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
}

func InitGrpcService() {
	interceptors := connect.WithInterceptors(
		validate.NewInterceptor(),
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

	gin.SetMode(gin.DebugMode)
	// routerGin := gin.Default()
	routerGin := c_.GetEngine()

	globalsMiddleware(routerGin)

	// routes ServiceHandler connectrpc
	routerGin.Any(path+"/*any", gin.WrapH(handler))
	// routes for anotations proto
	routerGin.Any("/productsapi/*any", gin.WrapH(transcoder))

	s := http.Server{
		Addr:      "localhost:8080",
		Handler:   routerGin,
		Protocols: c_.CreateProtoHTTP2NotTLS(),
	}

	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
