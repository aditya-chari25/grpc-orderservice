package inventory

import (
	"context"
	// "log"

	inventorypb "lfdtassgn/pkg/proto/inventory/proto"
)

type InventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	products map[int32]*inventorypb.Product
}

func NewInventoryServer() *InventoryServer {
	return &InventoryServer{
		products: map[int32]*inventorypb.Product{
			1: {Id: 1, Name: "Keyboard", Stock: 10, Price: 499.99},
			2: {Id: 2, Name: "Mouse", Stock: 25, Price: 299.99},
		},
	}
}

func (s *InventoryServer) CheckStock(ctx context.Context, req *inventorypb.CheckStockRequest) (*inventorypb.CheckStockResponse, error) {
	product, ok := s.products[req.ProductId]
	if !ok {
		return &inventorypb.CheckStockResponse{Available: false, Stock: 0}, nil
	}
	return &inventorypb.CheckStockResponse{
		Available: product.Stock >= req.Quantity,
		Stock:     product.Stock,
	}, nil
}

func (s *InventoryServer) UpdateStock(ctx context.Context, req *inventorypb.UpdateStockRequest) (*inventorypb.UpdateStockResponse, error) {
	product, ok := s.products[req.ProductId]
	if !ok {
		return &inventorypb.UpdateStockResponse{Success: false, Message: "Product not found"}, nil
	}
	if product.Stock < req.Quantity {
		return &inventorypb.UpdateStockResponse{Success: false, Message: "Insufficient stock"}, nil
	}
	product.Stock -= req.Quantity
	return &inventorypb.UpdateStockResponse{Success: true, Message: "Stock updated"}, nil
}

func (s *InventoryServer) GetProduct(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.GetProductResponse, error) {
	product, ok := s.products[req.ProductId]
	if !ok {
		return nil, nil
	}
	return &inventorypb.GetProductResponse{Product: product}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	var products []*inventorypb.Product
	for _, p := range s.products {
		products = append(products, p)
	}
	return &inventorypb.ListProductsResponse{Products: products, TotalCount: int32(len(products))}, nil
}
