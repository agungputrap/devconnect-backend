package http

import (
	"github.com/agungputrap/devconnect-backend/internal/application/user/dto"
	"github.com/agungputrap/devconnect-backend/internal/application/user/usecases"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	registerUserCase *usecases.RegisterUseCase
	loginUserCase    *usecases.LoginUseCase
}

func NewUserHandler(app fiber.Router, registerUC *usecases.RegisterUseCase, loginUC *usecases.LoginUseCase) {
	handler := &UserHandler{
		registerUserCase: registerUC,
		loginUserCase:    loginUC,
	}

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.registerUserCase.Execute(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.loginUserCase.Execute(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}
