package user

import (
	"context"

	userpb "lfdtassgn/pkg/proto/user/proto"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	users map[int32]*userpb.User
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: map[int32]*userpb.User{
			1: {Id: 1, Name: "Alice", Email: "alice@example.com"},
			2: {Id: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, ok := s.users[req.Id]
	if !ok {
		return nil, nil
	}
	return &userpb.GetUserResponse{User: user}, nil
}
