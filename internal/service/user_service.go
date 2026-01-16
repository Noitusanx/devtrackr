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

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Register - Complete registration flow
func (s *UserService) Register(ctx context.Context, name, email, password string) (*model.User, string, error) {
    // 1. Check if email already exists
    existingUser, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return nil, "", err
    }
    if existingUser != nil {
        return nil, "", util.ErrConflict("email already registered")
    }

    // 2. Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, "", util.ErrInternalServer("failed to process password")
    }

    // 3. Create user
    user := &model.User{
        ID:           uuid.New(),
        Name:         name,
        Email:        email,
        PasswordHash: string(hashedPassword),
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    if err := s.repo.Create(ctx, user); err != nil {
        return nil, "", err
    }

    // 4. Generate JWT token
    token, err := util.GenerateJWT(user.ID.String(), user.Email)
    if err != nil {
        return nil, "", util.ErrInternalServer("failed to generate token")
    }

    return user, token, nil
}

// Login - Complete authentication flow
func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, string, error) {
    // 1. Find user by email
    user, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return nil, "", util.ErrUnauthorized("invalid email or password")
    }
    if user == nil {
        return nil, "", util.ErrUnauthorized("invalid email or password")
    }

    // 2. Validate password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, "", util.ErrUnauthorized("invalid email or password")
    }

    // 3. Generate JWT token
    token, err := util.GenerateJWT(user.ID.String(), user.Email)
    if err != nil {
        return nil, "", util.ErrInternalServer("failed to generate token")
    }

    return user, token, nil
}

// GetCurrentUser - Get authenticated user
func (s *UserService) GetCurrentUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, util.ErrNotFound("user not found")
    }

    // Remove password hash from response
    user.PasswordHash = ""
    return user, nil
}

// UpdateProfile - Update user profile
func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, name string) (*model.User, error) {
    // 1. Get existing user
    user, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, util.ErrNotFound("user not found")
    }

    // 2. Update fields
    user.Name = name
    user.UpdatedAt = time.Now()

    // 3. Save
    if err := s.repo.Update(ctx, user); err != nil {
        return nil, err
    }

    // Remove password hash
    user.PasswordHash = ""
    return user, nil
}

// DeleteAccount - Delete user account
func (s *UserService) DeleteAccount(ctx context.Context, id uuid.UUID) error {
    // 1. Check if user exists
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return err
    }
    if user == nil {
        return util.ErrNotFound("user not found")
    }

    // 2. Delete user
    if err := s.repo.Delete(ctx, id); err != nil {
        return err
    }

    return nil
}