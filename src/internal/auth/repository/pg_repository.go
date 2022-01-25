package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

// authRepository - репозиторий авторизации.
type authRepository struct {
	db     *pgxpool.Pool
	logger *logg.Logg
}

// NewAuthRepository - конструктор репозитория авторизации.
func NewAuthRepository(db *pgxpool.Pool, logger *logg.Logg) *authRepository {
	return &authRepository{db: db, logger: logger}
}

func (r *authRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.Id = uuid.New()
	u := &entity.User{}
	if err := r.db.QueryRow(ctx, createUserQuery, &user.Id, &user.Login, &user.Password).Scan(u); err != nil {
		r.logger.LogError("authRepository.Create.Scan", err.Error())
		return nil, err
	}
	return u, nil
}

func (r *authRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := &entity.User{}
	if _, err := r.db.Query(ctx, updateUserQuery, u, &user.Login, &user.Password, &user.Id); err != nil {
		r.logger.LogError("authRepository.Update.Query", err.Error())
		return nil, err
	}
	return u, nil
}

func (r *authRepository) Delete(ctx context.Context, userId uuid.UUID) error {

	return nil
}

func (r *authRepository) GetById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	return nil, nil
}

func (r *authRepository) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {
	return nil, nil
}

func (r *authRepository) FindAll(ctx context.Context) (*entity.UsersList, error) {
	rows, err := r.db.Query(ctx, findAllUsers)
	if err != nil {
		return nil, err
	}

	var users = make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			r.logger.LogError("", err.Error())
			return nil, err
		}
		users = append(users, &user)
	}
	defer r.db.Close()

	return &entity.UsersList{Users: users}, nil
}
