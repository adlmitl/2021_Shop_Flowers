package auth

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/entity"
)

// Repository - Repository user.
type Repository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	// Delete UpdateUserLogin(ctx context.Context, userId uuid.UUID, login string) error
	Delete(ctx context.Context, userId uuid.UUID) error
	FindById(ctx context.Context, userId uuid.UUID) (*entity.User, error)
	FindByLogin(ctx context.Context, userLogin string) (*entity.User, error)
	FindAll(ctx context.Context) (*entity.UsersList, error)
}
