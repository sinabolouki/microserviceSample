package main

import (
	"context"
	"fmt"
	"sync"

	cataloguepb "microservice-sample/catalogue-service/gen"
)

type CatalogueServer struct {
	cataloguepb.UnimplementedCatalogueServiceServer
	mu     sync.Mutex
	items  map[string]*cataloguepb.Item
	nextID int
}

func NewCatalogueServer() *CatalogueServer {
	return &CatalogueServer{
		items: make(map[string]*cataloguepb.Item),
	}
}

func (s *CatalogueServer) CreateItem(ctx context.Context, req *cataloguepb.CreateItemRequest) (*cataloguepb.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	id := fmt.Sprintf("%d", s.nextID)
	item := &cataloguepb.Item{
		Id:    id,
		Title: req.Title,
		Uom:   req.Uom,
	}
	s.items[id] = item
	return item, nil
}

func (s *CatalogueServer) GetItem(ctx context.Context, req *cataloguepb.GetItemRequest) (*cataloguepb.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[req.Id]
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	return item, nil
}

func (s *CatalogueServer) ListItems(ctx context.Context, _ *cataloguepb.Empty) (*cataloguepb.ItemList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var itemList []*cataloguepb.Item
	for _, item := range s.items {
		itemList = append(itemList, item)
	}
	return &cataloguepb.ItemList{Items: itemList}, nil
}

func (s *CatalogueServer) ValidateItems(ctx context.Context, req *cataloguepb.ValidateItemsRequest) (*cataloguepb.ValidateItemsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var missing []string
	for _, id := range req.Ids {
		if _, ok := s.items[id]; !ok {
			missing = append(missing, id)
		}
	}
	return &cataloguepb.ValidateItemsResponse{
		AllFound:   len(missing) == 0,
		MissingIds: missing,
	}, nil
}
