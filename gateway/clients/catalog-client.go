package clients

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	userpb "microservice-sample/catalogue-service/gen"

	"google.golang.org/grpc"
)

func NewCatalogClient(address string) userpb.CatalogueServiceClient {
	conn, user_error := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if user_error != nil {
		log.Fatalf("Failed to connect to UserService: %v", user_error)
	}
	return userpb.NewCatalogueServiceClient(conn)
}
