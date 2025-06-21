package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"google.golang.org/grpc"

	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen"
)

func main() {
	// Connect to User and Catalogue Services
	userConn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	catConn, _ := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	userClient := userpb.NewUserServiceClient(userConn)
	catClient := cataloguepb.NewCatalogueServiceClient(catConn)

	lis, _ := net.Listen("tcp", ":50053")
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, NewOrderServer(catClient, userClient))

	log.Println("✅ OrderService running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
