package inventory

import (
	"context"
	"time"

	pbinv "lfdtassgn/pkg/proto/inventory/proto"

	"google.golang.org/grpc"
)

type InventoryClient struct {
	Client pbinv.InventoryServiceClient
}

func NewInventoryClient(addr string) (*InventoryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pbinv.NewInventoryServiceClient(conn)
	return &InventoryClient{Client: client}, nil
}

func (c *InventoryClient) CheckStock(productID, quantity int32) (*pbinv.CheckStockResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.Client.CheckStock(ctx, &pbinv.CheckStockRequest{
		ProductId: productID,
		Quantity:  quantity,
	})
}

func (c *InventoryClient) UpdateStock(productID, quantity int32) (*pbinv.UpdateStockResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return c.Client.UpdateStock(ctx, &pbinv.UpdateStockRequest{
		ProductId: productID,
		Quantity:  quantity,
	})
}
