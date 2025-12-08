package controller

import (
	"unc/services/unc-service/application/service"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	AuthController AuthController
}

func SetupControllers(services *service.Services, validator *validator.Validate) *Controllers {
	return &Controllers{
		AuthController: NewAuthController(services.AuthService, validator),
	}
}
