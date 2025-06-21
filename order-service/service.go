package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	cataloguepb "microservice-sample/catalogue-service/gen"
	orderpb "microservice-sample/order-service/gen"
	userpb "microservice-sample/user-service/gen"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	CatalogueSvc cataloguepb.CatalogueServiceClient
	UserSvc      userpb.UserServiceClient
	DB           *sql.DB
}

func NewOrderServer(c cataloguepb.CatalogueServiceClient, u userpb.UserServiceClient, db *sql.DB) *OrderServer {
	return &OrderServer{
		CatalogueSvc: c,
		UserSvc:      u,
		DB:           db,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.Order, error) {
	// Validate user
	_, err := s.UserSvc.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Validate catalogue items
	var itemIDs []string
	for _, pos := range req.Positions {
		itemIDs = append(itemIDs, pos.CatalogueItemId)
	}
	validateResp, err := s.CatalogueSvc.ValidateItems(ctx, &cataloguepb.ValidateItemsRequest{Ids: itemIDs})
	if err != nil || !validateResp.AllFound {
		return nil, fmt.Errorf("invalid catalogue items: %v", err)
	}

	// Begin DB transaction
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	orderID := uuid.NewString()
	_, err = tx.ExecContext(ctx, `INSERT INTO orders (id, user_id) VALUES ($1, $2)`, orderID, req.UserId)
	if err != nil {
		return nil, err
	}

	var positions []*orderpb.OrderPosition
	for _, pos := range req.Positions {
		posID := uuid.NewString()
		positions = append(positions, &orderpb.OrderPosition{
			Id:              posID,
			CatalogueItemId: pos.CatalogueItemId,
			Title:           pos.Title,
			Quantity:        pos.Quantity,
		})
		_, err := tx.ExecContext(ctx,
			`INSERT INTO order_positions (id, order_id, catalogue_item_id, title, quantity)
			 VALUES ($1, $2, $3, $4, $5)`,
			posID, orderID, pos.CatalogueItemId, pos.Title, pos.Quantity)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &orderpb.Order{
		Id:        orderID,
		UserId:    req.UserId,
		Positions: positions,
	}, nil
}

func (s *OrderServer) ListOrders(ctx context.Context, _ *orderpb.Empty) (*orderpb.OrderList, error) {
	rows, err := s.DB.QueryContext(ctx, `
		SELECT o.id, o.user_id,
		       op.id, op.catalogue_item_id, op.title, op.quantity
		FROM orders o
		JOIN order_positions op ON o.id = op.order_id
		ORDER BY o.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderMap := make(map[string]*orderpb.Order)
	for rows.Next() {
		var (
			orderID, userID, posID, itemID, title string
			quantity                              int32
		)
		if err := rows.Scan(&orderID, &userID, &posID, &itemID, &title, &quantity); err != nil {
			return nil, err
		}

		order, exists := orderMap[orderID]
		if !exists {
			order = &orderpb.Order{
				Id:        orderID,
				UserId:    userID,
				Positions: []*orderpb.OrderPosition{},
			}
			orderMap[orderID] = order
		}

		order.Positions = append(order.Positions, &orderpb.OrderPosition{
			Id:              posID,
			CatalogueItemId: itemID,
			Title:           title,
			Quantity:        quantity,
		})
	}

	var orders []*orderpb.Order
	for _, o := range orderMap {
		orders = append(orders, o)
	}

	return &orderpb.OrderList{Orders: orders}, nil
}
