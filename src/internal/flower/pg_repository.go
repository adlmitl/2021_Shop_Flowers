package flower

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/entity"
)

// Repository - Repository flower.
type Repository interface {
	Create(ctx context.Context, flower *entity.Flower) (*entity.Flower, error)
	Update(ctx context.Context, flower *entity.Flower) (*entity.Flower, error)
	Delete(ctx context.Context, flowerId uuid.UUID) error
	FindById(ctx context.Context, flowerId uuid.UUID) (*entity.Flower, error)
	FindByName(ctx context.Context, flowerName string) (*entity.Flower, error)
	FindAll(ctx context.Context) (*entity.FlowerList, error)
}
