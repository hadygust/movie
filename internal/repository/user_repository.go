package repository

import (
	"github.com/hadygust/movie/internal/dto"
	"github.com/hadygust/movie/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo *UserRepository) AllUser() ([]model.User, error) {
	var users []model.User
	err := repo.db.Find(&users).Error

	return users, err
}

func (repo *UserRepository) CreateUser(user model.User) (model.User, error) {
	res := repo.db.Create(user)
	err := res.Error

	return user, err
}

func (repo *UserRepository) LoginUserEmail(input dto.LoginRequest) (model.User, error) {
	var user model.User
	err := repo.db.Where("email = ?", input.Email).Find(&user).Error

	return user, err
}
