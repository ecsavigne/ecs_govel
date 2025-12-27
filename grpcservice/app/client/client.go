package main

import (
	"context"
	pbProduct "ecs_govel/grpcservice/gen/productpb/v1"
	pbService "ecs_govel/grpcservice/gen/services/v1"
	conn "ecs_govel/grpcservice/gen/services/v1/servicesv1connect"
	"fmt"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

func createProductServer(cl conn.ProductServiceClient, pR *connect.Request[pbService.CreateProductRequest]) *connect.Response[pbService.CreateProductResponse] {
	resp, err := cl.CreateProduct(context.Background(), pR)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func getProducts(cl conn.ProductServiceClient, pR *connect.Request[emptypb.Empty]) *connect.Response[pbService.GetProductsResponse] {
	pR.Header().Set("X-User-ID", "1342342342423")
	resp, err := cl.GetProducts(context.Background(), pR)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func createProductClient() conn.ProductServiceClient {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	client := conn.NewProductServiceClient(
		&http.Client{Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			Protocols:           p,
		}},
		"http://localhost:8080",
		connect.WithGRPC(),
		// connect.WithHTTPGet(),
	)

	return client
}

func main() {
	fmt.Println("Init Client product")

	cl := createProductClient()

	p := &pbProduct.Product{}
	p.SetId("1")
	p.SetName("Edilberto Coello Savigne")
	p.SetDescription("Edilberto Coello Savigne")
	p.SetPrice(20000)

	req := &pbService.CreateProductRequest{}
	req.SetProduct(p)

	fmt.Println("Call RPC with, 'ProductRequest' = ", req.GetProduct())

	// resp := createProductServer(cl, connect.NewRequest(req))
	resp := getProducts(cl, connect.NewRequest(&emptypb.Empty{}))
	fmt.Println("Response from derver RPC (ProductResponse): ", resp.Msg)
}
