package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Client *UserClient
}

func NewUserHandler(client *UserClient) *UserHandler {
	return &UserHandler{Client: client}
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	resp, err := h.Client.GetUser(int32(userID))
	if err != nil {
		http.Error(w, "User service error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
