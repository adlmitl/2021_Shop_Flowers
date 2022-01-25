package http

import (
	"context"
	"encoding/json"
	"net/http"
	"shopflowers/src/internal/auth"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

type AuthHandler struct {
	authService auth.Service
	l           *logg.Logg
}

func NewAuthHandler(authService auth.Service, l *logg.Logg) *AuthHandler {
	return &AuthHandler{authService: authService, l: l}
}

func (h *AuthHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.FindAll)
	mux.HandleFunc("/create", h.Create)
}

func (h *AuthHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.authService.FindAll(context.TODO())
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(users); err != nil {
		h.l.LogError("ErrExecTmplParse", err.Error())
	}
}

func (h *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	u := &entity.User{}
	user, err := h.authService.Create(context.TODO(), u)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		h.l.LogError("ErrExecTmplParse", err.Error())
	}
}
