package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.Register(r.Context(), req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	tokens, err := h.Service.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func (h *Handler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	tokens, err := h.Service.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	var req struct {
        RefreshToken string `json:"refresh_token"`
    }
	json.NewDecoder(r.Body).Decode(&req)

    err := h.Service.Logout(r.Context(), req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
    w.WriteHeader(http.StatusNoContent)
}