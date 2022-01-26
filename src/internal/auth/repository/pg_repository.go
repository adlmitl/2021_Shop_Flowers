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
	db        *pgxpool.Pool
	newLogger *logg.CommonLogger
}

// NewAuthRepository - конструктор репозитория авторизации.
func NewAuthRepository(db *pgxpool.Pool, newLogger *logg.CommonLogger) *authRepository {
	return &authRepository{db: db, newLogger: newLogger}
}

func (r *authRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.Id = uuid.New()
	var u entity.User
	if err := r.db.QueryRow(
		ctx,
		createUserQuery,
		user.Id,
		user.Login,
		user.Password,
	).Scan(&u.Id); err != nil {
		r.newLogger.Error("authRepository.Create.Scan", err.Error())
		return nil, err
	}

	return user, nil
}

func (r *authRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	var u entity.User
	if err := r.db.QueryRow(
		ctx,
		updateUserQuery,
		user.Login,
		user.Password,
		user.Id,
	).Scan(
		&u.Id,
		&u.Login,
		&u.Password,
	); err != nil {
		r.newLogger.Error("authRepository.Update.Query", err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *authRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	result, err := r.db.Exec(ctx, deleteByIdQuery, userId)
	if err != nil {
		r.newLogger.Error("db.Exec.Delete", err.Error())
		return err
	}
	if result.RowsAffected() == 0 {
		r.newLogger.Error("RowsAffected", err.Error())
		return err
	}
	return nil
}

func (r *authRepository) GetById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {

	var u entity.User
	if err := r.db.QueryRow(ctx, getUserByIdQuery, userId).Scan(
		&u.Id,
		&u.Login,
		&u.Password,
	); err != nil {
		r.newLogger.Error("db.QueryRow.GetById.Scan", err.Error())
		return nil, err
	}
	return &u, nil
}

func (r *authRepository) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {
	var u entity.User
	if err := r.db.QueryRow(ctx, getUserByLoginQuery, userLogin).Scan(
		&u.Id,
		&u.Login,
		&u.Password,
	); err != nil {
		r.newLogger.Error("db.QueryRow.FindByLogin.Scan", err.Error())
		return nil, err
	}
	return &u, nil
}

func (r *authRepository) FindAll(ctx context.Context) (*entity.UsersList, error) {
	rows, err := r.db.Query(ctx, findAllUsers)
	if err != nil {
		r.newLogger.Error("db.Query.FindAll", err.Error())
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
		); err != nil {
			r.newLogger.Error("rows.FindAll.Scan", err.Error())
			return nil, err
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		r.newLogger.Error("rows.Err.FindAll", err.Error())
		return nil, err
	}

	return &entity.UsersList{Users: users}, nil
}
