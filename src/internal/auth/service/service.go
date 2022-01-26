package service

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/auth"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

type authService struct {
	authRepo  auth.Repository
	newLogger *logg.CommonLogger
}

func NewAuthService(authRepo auth.Repository, newLogger *logg.CommonLogger) *authService {
	return &authService{authRepo: authRepo, newLogger: newLogger}
}

func (u *authService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {

	return u.authRepo.Create(ctx, user)
}

func (u *authService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {

	return u.authRepo.Update(ctx, user)
}

func (u *authService) Delete(ctx context.Context, userId uuid.UUID) error {

	return u.authRepo.Delete(ctx, userId)
}

func (u *authService) GetById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {

	return u.authRepo.GetById(ctx, userId)
}

func (u *authService) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {

	return u.authRepo.FindByLogin(ctx, userLogin)
}

func (u *authService) FindAll(ctx context.Context) (*entity.UsersList, error) {

	return u.authRepo.FindAll(ctx)
}
