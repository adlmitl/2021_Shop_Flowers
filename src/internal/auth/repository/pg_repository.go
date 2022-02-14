package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

// authRepository - Authentication.
type authRepository struct {
	db        *pgxpool.Pool
	newLogger *logg.CommonLogger
}

// NewAuthRepository - Constructor.
func NewAuthRepository(db *pgxpool.Pool, newLogger *logg.CommonLogger) *authRepository {
	return &authRepository{db: db, newLogger: newLogger}
}

// Create - Create user.
func (r *authRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	defer r.db.Close()

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

// Update - Update user.
func (r *authRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	defer r.db.Close()

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

// Delete - Delete user.
func (r *authRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	defer r.db.Close()

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

// FindById - Find user by id.
func (r *authRepository) FindById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	defer r.db.Close()

	var u entity.User
	if err := r.db.QueryRow(ctx, getUserByIdQuery, userId).Scan(
		&u.Id,
		&u.Login,
		&u.Password,
	); err != nil {
		r.newLogger.Error("db.QueryRow.FindById.Scan", err.Error())
		return nil, err
	}
	return &u, nil
}

// FindByLogin - Find user by login.
func (r *authRepository) FindByLogin(ctx context.Context, userLogin string) (*entity.User, error) {
	defer r.db.Close()

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

// FindAll - Find all users.
func (r *authRepository) FindAll(ctx context.Context) (*entity.UsersList, error) {
	defer r.db.Close()

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
