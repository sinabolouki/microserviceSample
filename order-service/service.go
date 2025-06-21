package main

import (
	"context"
	"fmt"
	"sync"

	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	mu           sync.Mutex
	orders       []*orderpb.Order
	nextOrderID  int
	nextPosID    int
	CatalogueSvc cataloguepb.CatalogueServiceClient
	UserSvc      userpb.UserServiceClient
}

func NewOrderServer(c cataloguepb.CatalogueServiceClient, u userpb.UserServiceClient) *OrderServer {
	return &OrderServer{
		CatalogueSvc: c,
		UserSvc:      u,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.Order, error) {
	// 1. Validate user
	_, err := s.UserSvc.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// 2. Validate catalogue items
	itemIDs := []string{}
	for _, pos := range req.Positions {
		itemIDs = append(itemIDs, pos.CatalogueItemId)
	}
	validateResp, err := s.CatalogueSvc.ValidateItems(ctx, &cataloguepb.ValidateItemsRequest{Ids: itemIDs})
	if err != nil || !validateResp.AllFound {
		return nil, fmt.Errorf("invalid catalogue items: %v", validateResp.MissingIds)
	}

	// 3. Create order
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextOrderID++
	orderID := fmt.Sprintf("%d", s.nextOrderID)

	var positions []*orderpb.OrderPosition
	for _, p := range req.Positions {
		s.nextPosID++
		posID := fmt.Sprintf("%d", s.nextPosID)
		positions = append(positions, &orderpb.OrderPosition{
			Id:              posID,
			CatalogueItemId: p.CatalogueItemId,
			Title:           p.Title,
			Quantity:        p.Quantity,
		})
	}

	order := &orderpb.Order{
		Id:        orderID,
		UserId:    req.UserId,
		Positions: positions,
	}
	s.orders = append(s.orders, order)
	return order, nil
}

func (s *OrderServer) ListOrders(ctx context.Context, _ *orderpb.Empty) (*orderpb.OrderList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &orderpb.OrderList{Orders: s.orders}, nil
}
