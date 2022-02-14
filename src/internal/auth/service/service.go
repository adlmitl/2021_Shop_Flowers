package service

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/auth"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

// authService - Authentication service user.
type authService struct {
	authRepo  auth.Repository
	newLogger *logg.CommonLogger
}

// NewAuthService - Constructor.
func NewAuthService(authRepo auth.Repository, newLogger *logg.CommonLogger) *authService {

	return &authService{authRepo: authRepo, newLogger: newLogger}
}

// Create - Create user.
func (u *authService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {

	return u.authRepo.Create(ctx, user)
}

// Update - Update user.
func (u *authService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {

	return u.authRepo.Update(ctx, user)
}

// Delete - Delete user by id.
func (u *authService) Delete(ctx context.Context, userId uuid.UUID) error {

	return u.authRepo.Delete(ctx, userId)
}

// FindById - Find user by id.
func (u *authService) FindById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {

	return u.authRepo.FindById(ctx, userId)
}

// FindByLogin - Find user by login.
func (u *authService) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {

	return u.authRepo.FindByLogin(ctx, userLogin)
}

// FindAll - Find all users.
func (u *authService) FindAll(ctx context.Context) (*entity.UsersList, error) {

	return u.authRepo.FindAll(ctx)
}
