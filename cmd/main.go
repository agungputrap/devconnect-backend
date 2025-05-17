package main

import (
	"github.com/agungputrap/devconnect-backend/internal/application/user/usecases"
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
	"github.com/agungputrap/devconnect-backend/internal/infrastructure/auth"
	"github.com/agungputrap/devconnect-backend/internal/infrastructure/config"
	"github.com/agungputrap/devconnect-backend/internal/infrastructure/persistence/postgres"
	"github.com/agungputrap/devconnect-backend/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	dbConfig := config.NewDatabaseConfig()
	dbProvider, err := postgres.NewDatabaseProvider(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	// Repositories
	userRepo := postgres.NewUserRepository(dbProvider.DB())

	// Infrastructure services
	passwordHasher := auth.NewBcryptHasher(0)

	// Services
	userService := user.NewUserService(passwordHasher)

	// Use cases
	registerUC := usecases.NewRegisterUseCase(userRepo, userService)
	loginUC := usecases.NewLoginUseCase(userRepo, userService)

	// HTTP handlers
	api := app.Group("/api")
	http.NewUserHandler(api.Group("/auth"), registerUC, loginUC)

	log.Fatal(app.Listen(":3000"))
}
