package user

import (
	"errors"
	"time"
)

type User struct {
	ID        uint     `gorm:"primary_key"`
	Name      string   `gorm:"type:varchar(255);not null"`
	Email     Email    `gorm:"type:varchar(255);unique;not null"`
	Password  Password `gorm:"column:password;type:varchar(255);not null"`
	Bio       string   `gorm:"type:text"`
	Skills    Skills   `gorm:"type:text"`
	GithubURL string   `gorm:"column:github_url;type:varchar(255)"`
	Role      Role     `gorm:"type:varchar(255);not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, email Email, password Password) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Role:      RoleUser, // Default role
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) UpdateProfile(bio string, skills Skills, githubURL string) {
	u.Bio = bio
	u.Skills = skills
	u.GithubURL = githubURL
	u.UpdatedAt = time.Now()
}

func (u *User) GetPassword() Password {
	return u.Password
}
