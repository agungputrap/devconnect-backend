package user

import "github.com/agungputrap/devconnect-backend/internal/domain/user"

type Repository interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}
