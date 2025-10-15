package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/govnocods/JWT-Authorization/utils"
)


func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCtxUser(r.Context())
	if user == nil {
		http.Error(w, "Unauthorized: user not found", http.StatusUnauthorized)
		return
	}

	response := map[string]any{
		"username": user.Username,
		"password": user.Password,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
