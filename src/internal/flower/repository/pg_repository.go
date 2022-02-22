package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"shopflowers/src/internal/entity"
	"shopflowers/src/pkg/logg"
)

// flowerRepository - flower repository.
type flowerRepository struct {
	db        *pgxpool.Pool
	newLogger *logg.CommonLogger
}

// NewFlowerRepository - Constructor.
func NewFlowerRepository(db *pgxpool.Pool, newLogger *logg.CommonLogger) *flowerRepository {
	return &flowerRepository{db: db, newLogger: newLogger}
}

// Create - Create flower.
func (r *flowerRepository) Create(ctx context.Context, flower *entity.Flower) (*entity.Flower, error) {
	flower.Id = uuid.New()
	var f entity.Flower
	if err := r.db.QueryRow(
		ctx,
		createFlowerQuery,
		flower.Id,
		flower.Name,
		flower.Price,
	).Scan(&f.Id); err != nil {
		r.newLogger.Error("flowerRepository.Create.Scan", err.Error())
		return nil, err
	}

	return flower, nil
}

// Update - Update flower.
func (r *flowerRepository) Update(ctx context.Context, flower *entity.Flower) (*entity.Flower, error) {

	var f entity.Flower
	if err := r.db.QueryRow(
		ctx,
		updateFlowerQuery,
		flower.Id,
		flower.Name,
		flower.Price,
	).Scan(
		&f.Id,
		&f.Name,
		&f.Price,
	); err != nil {
		r.newLogger.Error("flowerRepository.Update.Scan", err.Error())
		return nil, err
	}

	return &f, nil
}

// Delete - Delete flower.
func (r *flowerRepository) Delete(ctx context.Context, flowerId uuid.UUID) error {
	result, err := r.db.Exec(ctx, deleteFlowerByIdQuery, flowerId)
	if err != nil {
		r.newLogger.Error("flowerRepository.db.Exec.Delete", err.Error())
		return err
	}
	if result.RowsAffected() == 0 {
		r.newLogger.Error("RowsAffected", err.Error())
		return err
	}

	return nil
}

// FindById - Find by id flower.
func (r *flowerRepository) FindById(ctx context.Context, flowerId uuid.UUID) (*entity.Flower, error) {
	var f entity.Flower
	if err := r.db.QueryRow(ctx, getFlowerByIdQuery, flowerId).Scan(
		&f.Id,
		&f.Name,
		&f.Price,
	); err != nil {
		r.newLogger.Error("flowerRepository.FindById.db.Scan", err.Error())
		return nil, err
	}

	return &f, nil
}

// FindByName - Find by name flower.
func (r *flowerRepository) FindByName(ctx context.Context, flowerName string) (*entity.Flower, error) {
	var f entity.Flower
	if err := r.db.QueryRow(ctx, getFlowerByNameQuery, flowerName).Scan(
		&f.Id,
		&f.Name,
		&f.Price,
	); err != nil {
		r.newLogger.Error("flowerRepository.FindByName.db.Scan", err.Error())
		return nil, err
	}

	return &f, nil
}

// FindAll - Find all flowers.
func (r *flowerRepository) FindAll(ctx context.Context) (*entity.FlowerList, error) {
	rows, err := r.db.Query(ctx, findAllFlowers)
	if err != nil {
		r.newLogger.Error("flowerRepository.FindAll.db.Query", err.Error())
		return nil, err
	}
	defer rows.Close()

	flowers := make([]*entity.Flower, 0)
	for rows.Next() {
		var flower entity.Flower
		if err = rows.Scan(
			&flower.Id,
			&flower.Name,
			&flower.Price,
		); err != nil {
			r.newLogger.Error("flowerRepository.FindAll.db.Scan", err.Error())
			return nil, err
		}
		flowers = append(flowers, &flower)
	}

	if err = rows.Err(); err != nil {
		r.newLogger.Error("flowerRepository.rows.Err.FindAll", err.Error())
		return nil, err
	}

	return &entity.FlowerList{Flowers: flowers}, nil
}
