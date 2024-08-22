package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "grpc-sample/proto"
)

type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    // Simulate creating a user and returning an ID
    id := "1234"
    return &pb.CreateUserResponse{Id: id}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    // Simulate retrieving a user
    return &pb.GetUserResponse{Name: "John Doe", Age: 30}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})

    log.Printf("Server is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
