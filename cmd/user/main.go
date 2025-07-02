package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "lfdtassgn/pkg/proto/user/proto"
	"lfdtassgn/internal/user"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := user.NewUserServer()
	userpb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("UserService is running on :50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
