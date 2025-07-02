package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	orderpb "lfdtassgn/pkg/proto/order/proto"
	"lfdtassgn/internal/order"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	inventoryClient := order.NewInventoryClient("localhost:50051")
	userClient := order.NewUserClient("localhost:50053")

	orderServer := order.NewOrderServer(inventoryClient, userClient)

	orderpb.RegisterOrderServiceServer(grpcServer, orderServer)

	log.Println("OrderService running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
