package service

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	repo repository.UserRepository
	jwt util.JWTProvider
}


func NewUserService(repo repository.UserRepository, jwt util.JWTProvider) *UserService {
	return &UserService{
		repo: repo,
		jwt:  jwt,
	}
}

// Register
func (s *UserService) Register(ctx context.Context, name, email, pass string) (*model.User, string, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err == nil && u != nil {
		return nil, "", util.ErrConflict("email already used")
	}


	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &model.User{
		ID:   uuid.New(),
		Name: name,
		Email: email,
		PasswordHash: string(hash),
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, "", err
	}
	token, err := s.jwt.Generate(user.ID)
	return user, token, err

}

// Login
func (s *UserService) Login(ctx context.Context, email, pass string) (*model.User, string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", util.ErrNotFound("user not found")
	}

	if user == nil {
		return nil, "", util.ErrNotFound("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pass)); err != nil {
		return nil, "", util.ErrUnauthorized("invalid credentials")
	}

	token, err := s.jwt.Generate(user.ID)
	return user, token, err
}