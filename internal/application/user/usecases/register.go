package usecases

import (
	"github.com/agungputrap/devconnect-backend/internal/application/user/dto"
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
)

type RegisterUseCase struct {
	userRepo    user.Repository
	userService user.UserService
}

func NewRegisterUseCase(repo user.Repository, service user.UserService) *RegisterUseCase {
	return &RegisterUseCase{userRepo: repo, userService: service}
}

func (uc *RegisterUseCase) Execute(req dto.RegisterUserRequest) error {
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return err
	}

	newUser, err := uc.userService.Register(req.Name, email, req.Password)
	if err != nil {
		return err
	}

	return uc.userRepo.Create(newUser)
}
