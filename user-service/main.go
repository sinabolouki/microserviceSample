package main

import (
	"context"
	"fmt"
	"log"
	"microservice-sample/config"
	"net"
	"strings"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	// Import the generated protobuf code

	pb "microservice-sample/user-service/gen"
)

type server struct {
	pb.UnimplementedUserServiceServer
	mu    sync.Mutex
	users map[string]*pb.UserResponse
}

func NewServer() *server {
	return &server{
		users: make(map[string]*pb.UserResponse),
	}
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.NewString()
	user := &pb.UserResponse{
		Id:    id,
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	s.users[id] = user
	return user, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.GetId()]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *server) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.UserResponse
	for _, user := range s.users {
		users = append(users, user)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func main() {
	_, port, found := strings.Cut(config.UserServiceAddress, ":")
	if !found {
		log.Fatalf("invalid address format: %s", config.CatalogueServiceAddress)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewServer())

	log.Println("User service running on", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
