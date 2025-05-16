package user

import (
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
	userRepo "github.com/agungputrap/devconnect-backend/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	Register(user *user.User) error
	Login(email, password string) (*user.User, error)
}

type usecase struct {
	repo userRepo.Repository
}

func NewUsecase(repo userRepo.Repository) Usecase {
	return &usecase{repo: repo}
}

func (uc usecase) Register(u *user.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return uc.repo.Create(u)
}

func (uc usecase) Login(email, password string) (*user.User, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
