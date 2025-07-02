package order

import (
	"context"
	"time"

	pborder "lfdtassgn/pkg/proto/order/proto"

	"google.golang.org/grpc"
)

type OrderClient struct {
	Client pborder.OrderServiceClient
}

func NewOrderClient(addr string) (*OrderClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pborder.NewOrderServiceClient(conn)
	return &OrderClient{Client: client}, nil
}

func (c *OrderClient) CreateOrder(userID int32, items []*pborder.OrderItem) (*pborder.CreateOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.Client.CreateOrder(ctx, &pborder.CreateOrderRequest{
		UserId: userID,
		Items:  items,
	})
}
