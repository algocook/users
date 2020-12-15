package methods

import (
	"context"

	pb "github.com/algocook/proto/users"
)

// UsersMainServer comment
type UsersMainServer struct{}

// GetUser gRPC function
func (s *UsersMainServer) GetUser(c context.Context, request *pb.GetUserRequest) (response *pb.GetUserResponse, err error) {
	response = &pb.GetUserResponse{}

	client, err := NewPostgresClient("users")
	if err != nil {
		return response, err
	}
	user, err := client.GetUserByID(request.Id)
	if err != nil {
		return response, err
	}

	response.Id = user.ID
	response.Username = user.Username
	response.Title = user.Title
	response.Description = user.Description
	return response, nil
}

// CheckUsername method
func (s *UsersMainServer) CheckUsername(c context.Context, request *pb.CheckUsernameRequest) (response *pb.CheckUsernameResponse, err error) {
	response = &pb.CheckUsernameResponse{}

	client, err := NewPostgresClient("users")
	if err != nil {
		return response, err
	}

	res, err := client.CheckUsernameAvialability(request.Username)
	if err != nil {
		return response, err
	}

	response.IsAvailable = res

	return response, nil
}

// PostUser gRPC function
func (s *UsersMainServer) PostUser(c context.Context, request *pb.PostUserRequest) (response *pb.PostUserResponse, err error) {
	response = &pb.PostUserResponse{}

	client, err := NewPostgresClient("users")
	if err != nil {
		return response, err
	}

	id, err := client.InsertNewUser(request.Username, request.Title, "")
	if err != nil {
		return response, err
	}

	response.Id = id

	return response, nil
}
