package user

type UserService interface {
	Register(name string, email Email, rawPassword string) (*User, error)
	ValidateCredentials(user *User, password string) error
}

type userService struct {
	passwordHasher PasswordHasher
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

func NewUserService(passwordHasher PasswordHasher) UserService {
	return &userService{passwordHasher: passwordHasher}
}

func (s *userService) Register(name string, email Email, rawPassword string) (*User, error) {
	hashedPassword, err := s.passwordHasher.Hash(rawPassword)
	if err != nil {
		return nil, err
	}

	password := NewPassword(hashedPassword)
	return NewUser(name, email, password)
}

func (s *userService) ValidateCredentials(user *User, password string) error {
	return s.passwordHasher.Compare(user.GetPassword().Hash(), password)
}
