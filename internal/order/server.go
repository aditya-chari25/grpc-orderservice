package order

import (
	"context"
	"log"
	"time"

	orderpb "lfdtassgn/pkg/proto/order/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	orders         map[int32]*orderpb.Order
	lastID         int32
	inventoryClient *InventoryClient
	userClient      *UserClient // 👈 new field
}


func NewOrderServer(inv *InventoryClient, user *UserClient) *OrderServer {
	return &OrderServer{
		orders:          make(map[int32]*orderpb.Order),
		lastID:          0,
		inventoryClient: inv,
		userClient:      user,
	}
}


func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	user, err := s.userClient.GetUser(req.UserId)
	if err != nil || user == nil {
		return &orderpb.CreateOrderResponse{
			Success: false,
			Message: "Invalid user",
		}, nil
	}

	log.Printf("Creating order for user: %s (ID %d)\n", user.Name, user.Id)

	for _, item := range req.Items {
		available, err := s.inventoryClient.CheckStock(item.ProductId, item.Quantity)
		if err != nil || !available {
			return &orderpb.CreateOrderResponse{
				Success: false,
				Message: "Product out of stock or inventory check failed",
			}, nil
		}
	}
	s.lastID++
	now := timestamppb.New(time.Now())

	total := float32(0)
	for _, item := range req.Items {
		total += item.TotalPrice
	}

	order := &orderpb.Order{
		Id:        s.lastID,
		UserId:    req.UserId,
		Items:     req.Items,
		Total:     total,
		Status:    "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.orders[s.lastID] = order

	log.Printf("Created order ID %d for user %d\n", s.lastID, req.UserId)

	return &orderpb.CreateOrderResponse{
		OrderId: s.lastID,
		Total:   total,
		Success: true,
		Message: "Order created",
		Status:  "pending",
	}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	order, ok := s.orders[req.OrderId]
	if !ok {
		return nil, nil
	}
	return &orderpb.GetOrderResponse{Order: order}, nil
}

func (s *OrderServer) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	var result []*orderpb.Order
	for _, o := range s.orders {
		if o.UserId == req.UserId {
			result = append(result, o)
		}
	}
	return &orderpb.ListOrdersResponse{Orders: result, TotalCount: int32(len(result))}, nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.UpdateOrderStatusResponse, error) {
	order, ok := s.orders[req.OrderId]
	if !ok {
		return &orderpb.UpdateOrderStatusResponse{Success: false, Message: "Order not found"}, nil
	}
	order.Status = req.Status
	order.UpdatedAt = timestamppb.New(time.Now())
	return &orderpb.UpdateOrderStatusResponse{Success: true, Message: "Order status updated"}, nil
}
