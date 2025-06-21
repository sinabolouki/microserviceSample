package graph

import (
	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen" // generated gRPC code
)

type Resolver struct {
	UserClient      userpb.UserServiceClient
	CatalogueClient cataloguepb.CatalogueServiceClient
	OrderClient     orderpb.OrderServiceClient
}
