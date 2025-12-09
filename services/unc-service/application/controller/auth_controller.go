package controller

import (
	"fmt"
	"unc/services/unc-service/application/service"
	"unc/services/unc-service/domain/request"

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
	var registerRequest request.RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := s.validator.StructCtx(c.Context(), &registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := s.authService.Register(c.Context(), &registerRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Created user %s", registerRequest.Username),
	})
}
