package controller

import (
	"unc/services/unc-service/application/service"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	AuthController     AuthController
	ClothingController ClothingController
}

func SetupControllers(services *service.Services, validator *validator.Validate) *Controllers {
	return &Controllers{
		AuthController:     NewAuthController(services.AuthService, validator),
		ClothingController: NewClothingController(services.ClothingService, validator),
	}
}
