package main

import (
	"algocook/users/pkg/methods"
	"net"

	pb "github.com/algocook/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &methods.UsersMainServer{}

	// Register gRPC server
	pb.RegisterUsersServer(s, srv)

	// Listern on 5300
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Start gRPC server
	if err := s.Serve(listener); err != nil {
		grpclog.Fatalf("failed on starting server: %v", err)
	}
}
