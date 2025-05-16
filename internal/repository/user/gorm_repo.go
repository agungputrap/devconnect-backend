package user

import (
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) Create(user *user.User) error {
	return r.DB.Create(user).Error
}

func (r *GormRepository) FindByEmail(email string) (*user.User, error) {
	var userByEmail user.User
	if err := r.DB.Where("email = ?", email).First(&userByEmail).Error; err != nil {
		return nil, err
	}
	return &userByEmail, nil
}
