package methods

import (
	"context"

	"github.com/algocook/proto/users"
)

// UsersMainServer comment
type UsersMainServer struct{}

// GetUser gRPC function
func (s *UsersMainServer) GetUser(c context.Context, request *users.GetUserRequest) (response *users.GetUserResponse, err error) {
	response = &users.GetUserResponse{}

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
func (s *UsersMainServer) CheckUsername(c context.Context, request *users.CheckUsernameRequest) (response *users.CheckUsernameResponse, err error) {
	response = &users.CheckUsernameResponse{}

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
func (s *UsersMainServer) PostUser(c context.Context, request *users.PostUserRequest) (response *users.PostUserResponse, err error) {
	response = &users.PostUserResponse{}

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
