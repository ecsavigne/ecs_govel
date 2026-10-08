package grpcserverinit

import (
	"context"
	services "ecs_govel/internal/grpcservice/gen/services/v1"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct{}

func (srv *ProductService) CreateProduct(ctx context.Context, req *services.CreateProductRequest) (*services.CreateProductResponse, error) {
	resp := &services.CreateProductResponse{}

	resp.SetProduct(req.GetProduct())
	msg := strings.Builder{}
	msg.WriteString("Product with name: ")
	msg.WriteString(req.GetProduct().GetName())
	msg.WriteString(" has been created")

	resp.SetMsg(msg.String())

	return resp, nil
}

func (srv *ProductService) GetProduct(ctx context.Context, reqWrapper *services.GetProductRequest) (*services.GetProductResponse, error) {

	return &services.GetProductResponse{}, nil
}

func (srv *ProductService) GetProducts(ctx context.Context, req *emptypb.Empty) (*services.GetProductsResponse, error) {
	// get header
	// fmt.Printf("X-User-ID: %s\n", req.Header().Get("X-User-ID"))

	return &services.GetProductsResponse{}, nil
}

func (srv *ProductService) UpdateProduct(ctx context.Context, reqWrapper *services.UpdateProductRequest) (*services.UpdateProductResponse, error) {

	return &services.UpdateProductResponse{}, nil
}

func (srv *ProductService) DeleteProduct(ctx context.Context, reqWrapper *services.DeleteProductRequest) (*services.DeleteProductResponse, error) {

	return &services.DeleteProductResponse{}, nil
}

func (srv *ProductService) DeleteProduct2(ctx context.Context, reqWrapper *services.DeleteProductRequest) (*services.DeleteProductResponse, error) {

	return &services.DeleteProductResponse{}, nil
}
