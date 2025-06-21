package main

import (
	"log"
	"microservice-sample/utils"
	"net"
	"strings"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "microservice-sample/catalogue-service/gen"
)

func main() {
	_, port, found := strings.Cut(utils.CatalogueServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", utils.CatalogueServiceAddress)
	}

	db := utils.InitDB()
	defer db.Close()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCatalogueServiceServer(grpcServer, NewCatalogueServer(db))

	log.Println("✅ CatalogueService running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
