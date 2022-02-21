package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"shopflowers/src/internal/auth"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

// AuthHandler - Authentication handler.
type AuthHandler struct {
	authService auth.Service
	newLogger   *logg.CommonLogger
}

// NewAuthHandler - Constructor.
func NewAuthHandler(authService auth.Service, newLogger *logg.CommonLogger) *AuthHandler {
	return &AuthHandler{authService: authService, newLogger: newLogger}
}

// FindAll - Find all users.
func (h *AuthHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.authService.FindAll(context.TODO())
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.authService.FindAll", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(users); err != nil {
		h.newLogger.Error("AuthHandler.FindAll.json.NewEncoder", err.Error())
		return
	}
}

// Create - Create user.
func (h *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	var us entity.User

	if err := json.NewDecoder(r.Body).Decode(&us); err != nil {
		h.newLogger.Error("AuthHandler.Create.json.NewDecoder", err.Error())
		return
	}

	user, err := h.authService.Create(context.TODO(), &us)
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.authService.Create", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(user); err != nil {
		h.newLogger.Error("AuthHandler.Create.json.NewEncoder", err.Error())
	}
}

// Update - Update user data.
func (h *AuthHandler) Update(w http.ResponseWriter, r *http.Request) {
	idUser, err := uuid.Parse(r.URL.Query().Get("idUser"))
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.Update.uuid.Parse", http.StatusInternalServerError, err.Error())
		return
	}

	var u entity.User
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.newLogger.Error("AuthHandler.Update.json.NewDecoder", err.Error())
		return
	}

	updateUser, err := h.authService.Update(context.TODO(), &entity.User{
		Id:       idUser,
		Login:    u.Login,
		Password: u.Password,
	})
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.authService.Update", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(updateUser); err != nil {
		h.newLogger.Error("AuthHandler.Update.json.NewEncoder", err.Error())
		return
	}
}

// Delete - Delete user.
func (h *AuthHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idUser, err := uuid.Parse(r.URL.Query().Get("idUser"))
	if err != nil {
		h.newLogger.Error("AuthHandler.Delete.uuid.Parse", err.Error())
		return
	}
	rowsDelete := h.authService.Delete(context.TODO(), idUser)

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(rowsDelete); err != nil {
		h.newLogger.Error("AuthHandler.Delete.json.NewEncoder", err.Error())
		return
	}
}

// FindById - Find user by id.
func (h *AuthHandler) FindById(w http.ResponseWriter, r *http.Request) {
	idUser, err := uuid.Parse(r.URL.Query().Get("idUser"))
	if err != nil {
		h.newLogger.Error("AuthHandler.FindById.uuid.Parse", err.Error())
		return
	}

	findUserById, err := h.authService.FindById(context.TODO(), idUser)
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.authService.FindById", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(findUserById); err != nil {
		h.newLogger.Error("AuthHandler.FindById.json.Encode", err.Error())
		return
	}
}

// FindByLogin - Find user by login.
func (h *AuthHandler) FindByLogin(w http.ResponseWriter, r *http.Request) {
	userLogin := r.URL.Query().Get("login")

	findByLoginUser, err := h.authService.FindByLogin(context.TODO(), userLogin)
	if err != nil {
		h.newLogger.ErrorResponse("AuthHandler.authService.FindByLogin", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(findByLoginUser); err != nil {
		h.newLogger.Error("AuthHandler.FindByLogin.json.NewEncoder", err.Error())
		return
	}
}
