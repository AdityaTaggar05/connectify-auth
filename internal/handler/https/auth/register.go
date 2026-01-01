package authhandler

import (
	"encoding/json"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.Register(r.Context(), req.Email, req.Password); err != nil {
		switch err {
			case authservice.ErrInvalidEmailFormat, authservice.ErrInvalidPasswordFormat:
				http.Error(w, err.Error(), http.StatusBadRequest)
			case authservice.ErrUserAlreadyExists:
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}
