package methods

import (
	"context"

	pb "github.com/algocook/proto/users"
)

// UsersMainServer comment
type UsersMainServer struct{}

// GetUser gRPC function
func (s *UsersMainServer) GetUser(c context.Context, request *pb.GetUserRequest) (response *pb.GetUserResponse, err error) {
	response = &pb.GetUserResponse{
		Id:       request.Id,
		Username: "username",
		Title:    "title",
	}
	return response, nil
}

// PostUser gRPC function
func (s *UsersMainServer) PostUser(c context.Context, request *pb.PostUserRequest) (response *pb.PostUserResponse, err error) {
	response = &pb.PostUserResponse{
		Id: 1,
	}
	return response, nil
}
