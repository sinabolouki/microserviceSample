package clients

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	userpb "microservice-sample/user-service/gen"

	"google.golang.org/grpc"
)

func NewUserClient(address string) userpb.UserServiceClient {
	conn, user_error := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if user_error != nil {
		log.Fatalf("Failed to connect to UserService: %v", user_error)
	}
	return userpb.NewUserServiceClient(conn)
}
