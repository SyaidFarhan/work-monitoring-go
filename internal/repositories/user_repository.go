package repositories

import (
	"context"

	"gorm.io/gorm"

	"work-monitoring/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) error {
	return repository.db.
		WithContext(ctx).
		Create(user).
		Error
}

func (repository *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := repository.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *userRepository) FindByID(
	ctx context.Context,
	id uint,
) (*models.User, error) {
	var user models.User

	err := repository.db.
		WithContext(ctx).
		First(&user, id).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
