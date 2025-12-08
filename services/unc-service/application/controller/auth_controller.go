package controller

import (
	"unc/services/unc-service/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authService service.AuthService
	validator   *validator.Validate
}

func NewAuthController(authService service.AuthService, validator *validator.Validate) AuthController {
	return &AuthControllerImpl{
		authService: authService,
		validator:   validator,
	}
}

func (s *AuthControllerImpl) Register(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON("")
}
