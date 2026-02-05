package controller

import (
	"errors"
	"fmt"
	"time"
	"unc/services/unc-service/application/helper"
	"unc/services/unc-service/application/service"
	"unc/services/unc-service/domain/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Verify(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
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
		"message": fmt.Sprintf("Check registered email for user %s", registerRequest.Username),
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

	token, refreshToken, err := c.authService.Login(ctx.Context(), &loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   false,
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       fmt.Sprintf("Login successful, Welcome %s", loginRequest.Username),
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (c *AuthControllerImpl) Verify(ctx *fiber.Ctx) error {
	var verifyRequest request.VerifyRequest

	if err := ctx.BodyParser(&verifyRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := c.validator.StructCtx(ctx.Context(), &verifyRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := c.authService.Verify(ctx.Context(), &verifyRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("User is now verified"),
	})
}

func (c *AuthControllerImpl) Refresh(ctx *fiber.Ctx) error {
	var refreshRequest request.RefreshRequest
	// fmt.Println(ctx.Cookies("token"))
	// fmt.Println(ctx.Cookies("refresh_token"))

	if err := ctx.BodyParser(&refreshRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// if err := c.validator.StructCtx(ctx.Context(), &refreshRequest); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	if refreshRequest.RefreshToken == "" {
		refreshRequest.RefreshToken = ctx.Cookies("refresh_token")
	}

	token, refreshToken, err := c.authService.Refresh(ctx.Context(), &refreshRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   false,
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       "token refreshed",
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (c *AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	if token == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": errors.New("invalid credentials"),
		})
	}
	user_id := helper.GetUserFromJwt(token)
	if err := c.authService.Logout(ctx.Context(), user_id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	cookie = fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "logout success",
	})

}
