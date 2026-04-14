package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hadygust/movie/internal/dto"
	"github.com/hadygust/movie/internal/env"
	"github.com/hadygust/movie/internal/model"
	"github.com/hadygust/movie/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredential = errors.New("invalid credentials")
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) AllUser() ([]dto.UserResponse, error) {
	users, err := s.repo.AllUser()

	if err != nil {
		return nil, err
	}

	var res []dto.UserResponse

	for _, user := range users {
		// fmt.Println(user.Name)
		res = append(res, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return res, nil
}

func (s *UserService) FindByID(id string) (dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s UserService) RegisterUser(user dto.RegisterRequest) (dto.UserResponse, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	real_user := model.User{
		ID:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: string(password),
	}

	res, err := s.repo.CreateUser(real_user)

	if err != nil {
		fmt.Println(err.Error())
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}, nil

}

func (s *UserService) Login(input dto.LoginRequest) (string, error) {
	user, err := s.repo.LoginUserEmail(input)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user.Email, err
	}

	key, err := env.GetSecret()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))

	return tokenString, err
}
