package postgres

import (
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email user.Email) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email.String()).First(&user).Error

	return &user, err
}

func (r *userRepository) UpdateProfile(user *user.User) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
