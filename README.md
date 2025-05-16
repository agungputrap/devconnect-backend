# DevConnect Backend

This is the backend service for **DevConnect**, a collaborative platform for developers to post and join open-source or side projects.

## 🚀 Tech Stack

- **Go** (Fiber framework)
- **GORM** ORM
- **PostgreSQL** for the database
- JWT for authentication
- Dotenv for config management

## 📁 Project Structure
```bash
├── cmd/                    # Entry point
│   └── main.go
├── config/                 # DB config
│   └── database.go
├── internal/               # Business logic
│   ├── domain/             # Entity definitions (core business rules)
│   │   ├── user/
│   │   │   └── user.go
│   │   ├── project/
│   │   │   └── project.go
│   │   └── application/
│   │       └── application.go
│   ├── usecase/            # Use cases (application logic)
│   │   ├── user/
│   │   │   └── usecase.go
│   │   └── project/
│   │       └── usecase.go
│   ├── repository/         # Repository interfaces and GORM implementations
│   │   ├── user/
│   │   │   ├── interface.go
│   │   │   └── gorm_repo.go
│   │   └── project/
│   │       └── ...
│   └── delivery/           # Handlers / HTTP delivery layer
│       ├── http/
│       │   ├── user_handler.go
│       │   └── project_handler.go
│       └── middleware/
│           └── auth.go
├── migration/              # Migration files
│   ├── 001_create_users.sql
│   └── ...
├── routes/                 # Fiber route groupings
│   └── routes.go
├── .env
├── .env.example            # Application configuration 
├── go.mod
└── README.md
```

## 🛠️ Setup & Run Locally
1. Clone the repo
```bash
git clone https://github.com/your-username/devconnect-backend.git
cd devconnect-backend
```
2. Install dependencies
```bash
go mod tidy
```
3. Copy `.env.example` to `.env`. Adjust configuration to development environment
```bash
cp .env.example .env
```
4. For database, I used `migrate`(https://github.com/golang-migrate/migrate). Run `migrate` command:
```bash
migrate -path migrations -database "postgres://db_username:db_password@localhost:5432/db_name?sslmode=disable" up
```
5Run application
```bash
go run cmd/main.gp
```
