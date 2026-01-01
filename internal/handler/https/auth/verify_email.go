package authhandler

import (
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
)

func (h *Handler) HandleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		http.Error(w, "token is required", http.StatusBadRequest)
		return
	}

	err := h.Service.VerifyEmail(r.Context(), token)

	if err != nil {
		switch err {
			case authservice.ErrInvalidToken:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte("email verified successfully"))
	w.WriteHeader(http.StatusOK)
}
