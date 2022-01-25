package service

import (
	"context"
	"github.com/google/uuid"
	"shopflowers/src/internal/auth"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

type authService struct {
	authRepo auth.Repository
	l        *logg.Logg
}

func (u *authService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.authRepo.Create(ctx, user)
}

func (u *authService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *authService) Delete(ctx context.Context, userId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u *authService) GetById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *authService) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(authRepo auth.Repository, l *logg.Logg) *authService {
	return &authService{authRepo: authRepo, l: l}
}

func (u *authService) FindAll(ctx context.Context) (*entity.UsersList, error) {

	return u.authRepo.FindAll(ctx)
}
