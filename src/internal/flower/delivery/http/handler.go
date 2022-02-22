package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"shopflowers/src/internal/entity"
	"shopflowers/src/internal/flower"
	"shopflowers/src/pkg/logg"
)

// FlowerHandler - Flower Handler.
type FlowerHandler struct {
	flowerService flower.Repository
	newLogger     *logg.CommonLogger
}

// NewFlowerHandler - Constructor.
func NewFlowerHandler(flowerService flower.Repository, newLogger *logg.CommonLogger) *FlowerHandler {
	return &FlowerHandler{flowerService: flowerService, newLogger: newLogger}
}

// FindAll - Find all flowers.
func (h *FlowerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	flowers, err := h.flowerService.FindAll(context.TODO())
	if err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.FindAll", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(flowers); err != nil {
		h.newLogger.Error("FlowerHandler.FindAll.json.NewEncoder", err.Error())
	}
}

// Create - Create flower.
func (h *FlowerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var f entity.Flower

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		h.newLogger.Error("FlowerHandler.Create.json.NewDecoder", err.Error())
		return
	}

	garden, err := h.flowerService.Create(context.TODO(), &f)
	if err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.Create.flowerService.Create", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(garden); err != nil {
		h.newLogger.Error("FlowerHandler.Create.json.NewEncoder", err.Error())
		return
	}
}

// Update - Update flower data.
func (h *FlowerHandler) Update(w http.ResponseWriter, r *http.Request) {
	idFlower, err := uuid.Parse(r.URL.Query().Get("idFlower"))
	if err != nil {
		h.newLogger.Error("FlowerHandler.Update.uuid.Parse", err.Error())
		return
	}
	var f entity.Flower

	if err = json.NewDecoder(r.Body).Decode(&f); err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.Update.json.NewDecoder", http.StatusInternalServerError, err.Error())
		return
	}

	updateFlower, err := h.flowerService.Update(context.TODO(), &entity.Flower{
		Id:    idFlower,
		Name:  f.Name,
		Price: f.Price,
	})
	if err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.flowerService.Update", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(updateFlower); err != nil {
		h.newLogger.Error("FlowerHandler.Update.json.NewEncoder", err.Error())
		return
	}
}

// Delete - Delete flower by id.
func (h *FlowerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idFlower, err := uuid.Parse(r.URL.Query().Get("idFlower"))
	if err != nil {
		h.newLogger.Error("FlowerHandler.Delete.uuid.Parse", err.Error())
		return
	}

	rowsDelete := h.flowerService.Delete(context.TODO(), idFlower)

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(rowsDelete); err != nil {
		h.newLogger.Error("FlowerHandler.Delete.json.NewEncoder", err.Error())
		return
	}
}

// FindById - Find flower by id.
func (h *FlowerHandler) FindById(w http.ResponseWriter, r *http.Request) {
	idFlower, err := uuid.Parse(r.URL.Query().Get("idFlower"))
	if err != nil {
		h.newLogger.Error("FlowerHandler.FindById.uuid.Parse", err.Error())
		return
	}

	findFlowerById, err := h.flowerService.FindById(context.TODO(), idFlower)
	if err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.flowerService.FindById", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(findFlowerById); err != nil {
		h.newLogger.Error("FlowerHandler.FindById.json.Encode", err.Error())
		return
	}
}

// FindByName - Find user by name.
func (h *FlowerHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	flowerLogin := r.URL.Query().Get("name")

	findByLoginFlower, err := h.flowerService.FindByName(context.TODO(), flowerLogin)
	if err != nil {
		h.newLogger.ErrorResponse("FlowerHandler.authService.FindByLogin", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(findByLoginFlower); err != nil {
		h.newLogger.Error("FlowerHandler.FindByName.json.NewEncoder", err.Error())
		return
	}
}
