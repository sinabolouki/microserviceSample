package main

import (
	"log"
	"microservice-sample/config"
	"net"
	"strings"

	"google.golang.org/grpc"
	cataloguepb "microservice-sample/catalogue-service/gen"
)

func main() {
	_, port, found := strings.Cut(config.CatalogueServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", config.CatalogueServiceAddress)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cataloguepb.RegisterCatalogueServiceServer(grpcServer, NewCatalogueServer())

	log.Println("✅ CatalogueService running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
