package user

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	Skills    string `json:"skills"`
	GithubURL string `json:"github_url"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
