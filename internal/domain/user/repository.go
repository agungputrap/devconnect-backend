package user

type Repository interface {
	Create(user *User) error
	FindByEmail(email Email) (*User, error)
	UpdateProfile(user *User)
	Delete(id uint) error
}
