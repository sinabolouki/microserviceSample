package graph

import (
	userpb "microservice-sample/user-service/gen" // generated gRPC code
)

type Resolver struct {
	UserClient userpb.UserServiceClient
}
