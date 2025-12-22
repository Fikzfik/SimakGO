package handler

import (
	"simak-go/app/helper"
	"simak-go/app/models"
	"simak-go/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Email == "" || req.Password == "" || req.Name == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "All fields are required")
	}

	user, err := h.authService.Register(c.Context(), req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, "Registration successful", user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	token, user, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return helper.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
		"user":  user,
	})
}
