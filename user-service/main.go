package main

import (
	"google.golang.org/grpc"
	"log"
	"microservice-sample/utils"
	"net"
	"strings"
	// Import the generated protobuf code

	pb "microservice-sample/user-service/gen"
)

func main() {
	_, port, found := strings.Cut(utils.UserServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", utils.UserServiceAddress)
	}

	db := utils.InitDB()
	defer db.Close()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewServer(db))

	log.Println("User service running on", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
