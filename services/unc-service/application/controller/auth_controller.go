package controller

import (
	"unc/services/unc-service/application/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (s *AuthControllerImpl) Register(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON("")
}
