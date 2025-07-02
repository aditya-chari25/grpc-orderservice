package main

import (
	"log"
	"net/http"

	inventorysvc "lfdtassgn/services/inventory"
	ordersvc "lfdtassgn/services/order"
	usersvc "lfdtassgn/services/user"
)

func main() {
	// Initialize gRPC clients
	inventoryClient, err := inventorysvc.NewInventoryClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to connect to InventoryService: %v", err)
	}

	orderClient, err := ordersvc.NewOrderClient("localhost:50052")
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}

	userClient, err := usersvc.NewUserClient("localhost:50053")
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}

	// Create HTTP handlers
	inventoryHandler := inventorysvc.NewInventoryHandler(inventoryClient)
	orderHandler := ordersvc.NewOrderHandler(orderClient)
	userHandler := usersvc.NewUserHandler(userClient)

	// Register routes
	http.HandleFunc("/check_stock", inventoryHandler.HandleCheckStock)
	http.HandleFunc("/update_stock", inventoryHandler.HandleUpdateStock)
	// http.HandleFunc("/get_product", inventoryHandler.HandleGetProduct)
	// http.HandleFunc("/list_products", inventoryHandler.HandleListProducts)

	http.HandleFunc("/create_order", orderHandler.HandleCreateOrder)
	// http.HandleFunc("/get_order", orderHandler.HandleGetOrder)
	// http.HandleFunc("/list_orders", orderHandler.HandleListOrders)
	// http.HandleFunc("/update_order_status", orderHandler.HandleUpdateOrderStatus)

	http.HandleFunc("/get_user", userHandler.HandleGetUser)

	// Start server
	log.Println("API Gateway running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
