package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"

	"github.com/google/uuid"

	pb "microservice-sample/catalogue-service/gen"
)

type CatalogueServer struct {
	pb.UnimplementedCatalogueServiceServer
	DB *sql.DB
}

func NewCatalogueServer(db *sql.DB) *CatalogueServer {
	return &CatalogueServer{DB: db}
}

func (s *CatalogueServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	id := uuid.NewString()
	_, err := s.DB.ExecContext(ctx,
		`INSERT INTO catalogue_items (id, title, uom) VALUES ($1, $2, $3)`,
		id, req.Title, req.Uom)
	if err != nil {
		return nil, fmt.Errorf("failed to insert catalogue item: %w", err)
	}

	return &pb.Item{Id: id, Title: req.Title, Uom: req.Uom}, nil
}

func (s *CatalogueServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.Item, error) {
	row := s.DB.QueryRowContext(ctx,
		`SELECT id, title, uom FROM catalogue_items WHERE id = $1`, req.Id)

	var item pb.Item
	if err := row.Scan(&item.Id, &item.Title, &item.Uom); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("catalogue item not found")
		}
		return nil, fmt.Errorf("failed to fetch item: %w", err)
	}
	return &item, nil
}

func (s *CatalogueServer) ListItems(ctx context.Context, _ *pb.Empty) (*pb.ItemList, error) {
	rows, err := s.DB.QueryContext(ctx,
		`SELECT id, title, uom FROM catalogue_items`)
	if err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}
	defer rows.Close()

	var items []*pb.Item
	for rows.Next() {
		var item pb.Item
		if err := rows.Scan(&item.Id, &item.Title, &item.Uom); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, &item)
	}

	return &pb.ItemList{Items: items}, nil
}

func (s *CatalogueServer) ValidateItems(ctx context.Context, req *pb.ValidateItemsRequest) (*pb.ValidateItemsResponse, error) {
	if len(req.Ids) == 0 {
		return &pb.ValidateItemsResponse{AllFound: true}, nil
	}

	query := `SELECT id FROM catalogue_items WHERE id = ANY($1)`
	rows, err := s.DB.QueryContext(ctx, query, pq.Array(req.Ids))
	if err != nil {
		return nil, fmt.Errorf("failed to validate items: %w", err)
	}
	defer rows.Close()

	found := map[string]bool{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		found[id] = true
	}

	var missing []string
	for _, id := range req.Ids {
		if !found[id] {
			missing = append(missing, id)
		}
	}

	return &pb.ValidateItemsResponse{
		AllFound:   len(missing) == 0,
		MissingIds: missing,
	}, nil
}
