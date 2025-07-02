package order

import (
	"context"
	"log"
	"time"

	userpb "lfdtassgn/pkg/proto/user/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}
	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (uc *UserClient) GetUser(id int32) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := uc.client.GetUser(ctx, &userpb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
