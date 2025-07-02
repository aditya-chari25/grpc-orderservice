package inventory

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type InventoryHandler struct {
	Client *InventoryClient
}

func NewInventoryHandler(client *InventoryClient) *InventoryHandler {
	return &InventoryHandler{Client: client}
}

func (h *InventoryHandler) HandleCheckStock(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
	quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))

	resp, err := h.Client.CheckStock(int32(productID), int32(quantity))
	if err != nil {
		http.Error(w, "Inventory check failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *InventoryHandler) HandleUpdateStock(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProductID int32 `json:"product_id"`
		Quantity  int32 `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.Client.UpdateStock(req.ProductID, req.Quantity)
	if err != nil {
		http.Error(w, "Inventory service error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

