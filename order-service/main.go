package main

import (
	"log"
	"microservice-sample/utils"
	"net"
	"strings"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen"
)

func main() {
	// gRPC clients
	userConn, _ := grpc.NewClient(utils.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	catConn, _ := grpc.NewClient(utils.CatalogueServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	userClient := userpb.NewUserServiceClient(userConn)
	catClient := cataloguepb.NewCatalogueServiceClient(catConn)

	// DB
	db := utils.InitDB()
	defer db.Close()

	// Port
	_, port, found := strings.Cut(utils.OrderServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", utils.OrderServiceAddress)
	}
	lis, _ := net.Listen("tcp", ":"+port)

	// Server
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, NewOrderServer(catClient, userClient, db))

	log.Println("✅ OrderService running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
