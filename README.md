# DevConnect Backend

This is the backend service for **DevConnect**, a collaborative platform for developers to post and join open-source or side projects.

## 🚀 Tech Stack

- **Go** (Fiber framework)
- **GORM** ORM
- **PostgreSQL** for the database
- JWT for authentication
- Dotenv for config management

## 📁 Project Structure
```
├── cmd/                          # Entry point
│   └── main.go
├── internal/                     # Business logic
│   ├── domain/                   # Entity definitions (core business rules)
│   │   ├── user/
│   │   │   └── user.go           # Domain model
│   │   │   └── repository.go     # Repository interface
│   │   │   └── service.go        # Domain service
│   │   │   └── value_objects.go  # Value objects
│   │   └── project/
│   │       └── ...
│   ├── application/            
│   │   └── user/
│   │       ├── usecases/
│   │       │   ├── register.go
│   │       │   └── login.go
│   │       └── dto/
│   │           └── user_dto.go
│   ├── infrastructure/
│   │   ├── config/                 # Configuration related
│   │   │   ├── config.go
│   │   │   └── database.go                
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       └── user_repository.go
│   │   └── auth/
│   │       └── jwt_service.go
│   └── interfaces/           
│       └── http/
│           └── user_handler.go
├── migration/                    # Migration files
│   ├── 001_create_users.sql
│   └── ...
├── .env
├── .env.example                  # Application configuration 
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
