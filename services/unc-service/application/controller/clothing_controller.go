package controller

import (
	"errors"
	"unc/services/unc-service/application/service"

	"unc/services/unc-service/application/helper"
	"unc/services/unc-service/domain/request"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ClothingController interface {
	Create(ctx *fiber.Ctx) error
}

type ClothingControllerImpl struct {
	clothingService service.ClothingService
	validator       *validator.Validate
}

func NewClothingController(clothingService service.ClothingService, validator *validator.Validate) ClothingController {
	return &ClothingControllerImpl{
		clothingService: clothingService,
		validator:       validator,
	}
}

func (c *ClothingControllerImpl) Create(ctx *fiber.Ctx) error {
	var createRequest request.CreateClothingRequest
	token := ctx.Locals("user").(*jwt.Token)
	if token == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": errors.New("invalid credentials"),
		})
	}
	user_id := helper.GetUserFromJwt(token)
	if err := ctx.BodyParser(&createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "image is required",
		})
	}
	createRequest.Image = image

	if err := c.validator.StructCtx(ctx.Context(), &createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.clothingService.CreateClothing(ctx.Context(), user_id, &createRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Clothing created successfully",
	})
}
