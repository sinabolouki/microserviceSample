package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"microservice-sample/config"
	"net"
	"strings"

	"google.golang.org/grpc"

	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen"
)

func main() {
	// Connect to User and Catalogue Services
	userConn, _ := grpc.NewClient(config.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	catConn, _ := grpc.NewClient(config.CatalogueServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	userClient := userpb.NewUserServiceClient(userConn)
	catClient := cataloguepb.NewCatalogueServiceClient(catConn)

	_, port, found := strings.Cut(config.OrderServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", config.CatalogueServiceAddress)
	}

	lis, _ := net.Listen("tcp", ":"+port)
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, NewOrderServer(catClient, userClient))

	log.Println("✅ OrderService running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
