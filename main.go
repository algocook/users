package main

import (
	"context"
	"net"

	pb "github.com/algocook/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUsersServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) GetUser(c context.Context, request *pb.GetUserRequest) (response *pb.GetUserResponse, err error) {
	response = &pb.GetUserResponse{
		Id:       request.Id,
		Username: "username",
		Title:    "title",
	}
	return response, nil
}

func (s *server) PostUser(c context.Context, request *pb.PostUserRequest) (response *pb.PostUserResponse, err error) {
	response = &pb.PostUserResponse{
		Id: 1,
	}
	return response, nil
}
