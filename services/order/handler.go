package order

import (
	"encoding/json"
	"net/http"

	pborder "lfdtassgn/pkg/proto/order/proto"
)

type OrderHandler struct {
	Client *OrderClient
}

func NewOrderHandler(client *OrderClient) *OrderHandler {
	return &OrderHandler{Client: client}
}

func (h *OrderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	type OrderItemRequest struct {
		ProductID  int32   `json:"product_id"`
		Quantity   int32   `json:"quantity"`
		UnitPrice  float32 `json:"unit_price"`
		TotalPrice float32 `json:"total_price"`
	}

	type Req struct {
		UserID int32              `json:"user_id"`
		Items  []OrderItemRequest `json:"items"`
	}

	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var orderItems []*pborder.OrderItem
	for _, item := range req.Items {
		orderItems = append(orderItems, &pborder.OrderItem{
			ProductId:  item.ProductID,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.TotalPrice,
		})
	}

	resp, err := h.Client.CreateOrder(req.UserID, orderItems)
	if err != nil {
		http.Error(w, "Order service error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
