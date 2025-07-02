package user

import (
	"context"
	"time"

	pbuser "lfdtassgn/pkg/proto/user/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client pbuser.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pbuser.NewUserServiceClient(conn)
	return &UserClient{Client: client}, nil
}

func (c *UserClient) GetUser(userID int32) (*pbuser.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.Client.GetUser(ctx, &pbuser.GetUserRequest{
		Id: userID,
	})
}
