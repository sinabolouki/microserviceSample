package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	cataloguepb "microservice-sample/catalogue-service/gen"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cataloguepb.RegisterCatalogueServiceServer(grpcServer, NewCatalogueServer())

	log.Println("✅ CatalogueService running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
