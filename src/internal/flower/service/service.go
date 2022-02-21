package service

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/entity"
	"shopflowers/src/internal/flower"
	"shopflowers/src/pkg/logg"
)

// flowerService - Service flowers.
type flowerService struct {
	flowerRepo flower.Repository
	newLogger  *logg.CommonLogger
}

// NewFlowerService - Constructor.
func NewFlowerService(flowerRepo flower.Repository, newLogger *logg.CommonLogger) *flowerService {
	return &flowerService{flowerRepo: flowerRepo, newLogger: newLogger}
}

// Create - Create flower.
func (f *flowerService) Create(ctx context.Context, flower *entity.Flower) (*entity.Flower, error) {

	return f.flowerRepo.Create(ctx, flower)
}

// Update - Update flower.
func (f *flowerService) Update(ctx context.Context, flower *entity.Flower) (*entity.Flower, error) {

	return f.flowerRepo.Update(ctx, flower)
}

// Delete - Delete flower.
func (f *flowerService) Delete(ctx context.Context, flowerId uuid.UUID) error {

	return f.flowerRepo.Delete(ctx, flowerId)
}

// FindById - Find by id flower.
func (f *flowerService) FindById(ctx context.Context, flowerId uuid.UUID) (*entity.Flower, error) {

	return f.flowerRepo.FindById(ctx, flowerId)
}

// FindByName - Find by name flower.
func (f *flowerService) FindByName(ctx context.Context, flowerName string) (*entity.Flower, error) {

	return f.flowerRepo.FindByName(ctx, flowerName)
}

// FindAll - Find all flowers.
func (f *flowerService) FindAll(ctx context.Context) (*entity.FlowerList, error) {

	return f.flowerRepo.FindAll(ctx)
}
