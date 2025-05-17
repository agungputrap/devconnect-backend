package dto

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Skills    string `json:"skills"`
	GithubURL string `json:"github_url"`
	Role      string `json:"role"`
}
