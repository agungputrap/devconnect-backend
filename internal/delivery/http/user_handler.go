package http

import (
	"github.com/agungputrap/devconnect-backend/internal/domain/user"
	uc "github.com/agungputrap/devconnect-backend/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase uc.Usecase
}

func NewUserHandler(app fiber.Router, usecase uc.Usecase) {
	handler := &UserHandler{usecase: usecase}

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var u user.User
	if err := c.BodyParser(&u); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.usecase.Register(&u); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{"message": "registration success"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userLogin, err := h.usecase.Login(input.Email, input.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"id":    userLogin.ID,
		"name":  userLogin.Name,
		"email": userLogin.Email,
	})
}
