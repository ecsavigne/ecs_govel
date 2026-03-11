package grpcserverinit

import (
	"context"
	services "ecs_govel/grpcservice/gen/services/v1"
	"fmt"
	"strings"

	"connectrpc.com/connect"
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

func (srv *ProductService) DeleteProduct2(ctx context.Context, reqWrapper *connect.Request[services.DeleteProductRequest]) (*connect.Response[services.DeleteProductResponse], error) {

	return connect.NewResponse(&services.DeleteProductResponse{}), nil
}
