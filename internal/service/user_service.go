package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hadygust/movie/internal/dto"
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
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}, err

}

func (s *UserService) Login(input dto.LoginRequest) (dto.UserResponse, error) {
	user, err := s.repo.LoginUserEmail(input)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, err
}
