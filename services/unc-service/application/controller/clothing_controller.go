package controller

import (
	"unc/services/unc-service/application/service"

	"github.com/go-playground/validator/v10"
)

type ClothingController interface {
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
