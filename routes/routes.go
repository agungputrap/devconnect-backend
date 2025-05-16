package routes

import (
	"github.com/agungputrap/devconnect-backend/config"
	"github.com/agungputrap/devconnect-backend/internal/delivery/http"
	"github.com/agungputrap/devconnect-backend/internal/repository/user"
	userUC "github.com/agungputrap/devconnect-backend/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := config.DB
	userRepo := user.NewGormRepository(db)
	userUseCase := userUC.NewUsecase(userRepo)

	api := app.Group("/api")
	http.NewUserHandler(api.Group("/auth"), userUseCase)
}
