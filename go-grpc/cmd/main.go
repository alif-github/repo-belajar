package main

import (
	"github.com/alif-github/go-grpc/cmd/config"
	"github.com/alif-github/go-grpc/cmd/services"
	productPb "github.com/alif-github/go-grpc/pb/product"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listened %v", err.Error())
	}

	//--- Connect DB
	config.SetServerAttribute()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err = grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
