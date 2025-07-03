package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	err := ur.database.Create(&user).Error
	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	err := ur.database.Find(&users).Error
	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.database.First(&user, "email = ?", email).Error
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id uint) (domain.User, error) {
	var user domain.User
	err := ur.database.First(&user, "id = ?", id).Error
	return user, err
}
