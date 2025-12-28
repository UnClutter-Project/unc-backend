package controller

import (
	"fmt"
	"unc/services/unc-service/application/service"
	"unc/services/unc-service/domain/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
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

func (c *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	var registerRequest request.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.validator.StructCtx(ctx.Context(), &registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.authService.Register(ctx.Context(), &registerRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Created user %s", registerRequest.Username),
	})
}

func (c *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.validator.StructCtx(ctx.Context(), &loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := c.authService.Login(ctx.Context(), &loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Login successful, Welcome %s", loginRequest.Username),
		"token":   token,
	})
}
