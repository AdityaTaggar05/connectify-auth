package authhandler

import (
	"encoding/json"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	tokens, err := h.Service.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		switch err {
			case authservice.ErrEmailNotVerified:
				http.Error(w, "email not verified", http.StatusForbidden)
				return
			default:
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}
