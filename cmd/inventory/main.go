package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	inventorypb "lfdtassgn/pkg/proto/inventory/proto"
	"lfdtassgn/internal/inventory"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	inventoryService := inventory.NewInventoryServer()
	inventorypb.RegisterInventoryServiceServer(grpcServer, inventoryService)

	log.Println("InventoryService is running on :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
