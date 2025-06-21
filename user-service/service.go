package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	pb "microservice-sample/user-service/gen"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{db: db}
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	id := uuid.NewString()
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
		id, req.GetName(), req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &pb.UserResponse{
		Id:    id,
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	row := s.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", req.GetId())
	var user pb.UserResponse
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}

func (s *server) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.ListUsersResponse, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*pb.UserResponse
	for rows.Next() {
		var user pb.UserResponse
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}
