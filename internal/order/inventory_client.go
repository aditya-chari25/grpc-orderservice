package order

import (
	"context"
	"log"
	"time"

	inventorypb "lfdtassgn/pkg/proto/inventory/proto"
	"google.golang.org/grpc"
)

type InventoryClient struct {
	client inventorypb.InventoryServiceClient
}

func NewInventoryClient(addr string) *InventoryClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to inventory service: %v", err)
	}
	return &InventoryClient{
		client: inventorypb.NewInventoryServiceClient(conn),
	}
}

func (ic *InventoryClient) CheckStock(productID, quantity int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := ic.client.CheckStock(ctx, &inventorypb.CheckStockRequest{
		ProductId: productID,
		Quantity:  quantity,
	})
	if err != nil {
		return false, err
	}
	return resp.Available, nil
}

func (ic *InventoryClient) UpdateStock(productID, quantity int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := ic.client.UpdateStock(ctx, &inventorypb.UpdateStockRequest{
		ProductId: productID,
		Quantity:  quantity,
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
