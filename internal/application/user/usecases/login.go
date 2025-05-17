package usecases

import (
	"github.com/agungputrap/devconnect-backend/internal/application/user/dto"
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
)

type LoginUseCase struct {
	userRepo    user.Repository
	userService user.UserService
}

func NewLoginUseCase(repo user.Repository, service user.UserService) *LoginUseCase {
	return &LoginUseCase{userRepo: repo, userService: service}
}

func (uc *LoginUseCase) Execute(req dto.LoginRequest) (*user.User, error) {
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := uc.userService.ValidateCredentials(user, req.Password); err != nil {
		return nil, err
	}

	return user, nil
}
