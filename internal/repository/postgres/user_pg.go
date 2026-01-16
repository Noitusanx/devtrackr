package postgres

import (
	"context"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) repository.UserRepository {
	return &UserPG{db}
}

func (r *UserPG) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// user_pg.go - FIXED
func (r *UserPG) FindByEmail(ctx context.Context, email string) (*model.User, error) {
    var user model.User
    err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
    
    // Handle record not found
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil  
        }
        return nil, err  
    }
    
    return &user, nil  
}

func (r *UserPG) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
    var user model.User
    err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
    
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    
    return &user, nil
}

func (r *UserPG) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserPG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}

