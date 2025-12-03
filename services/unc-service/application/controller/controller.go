package controller

import (
	"unc/services/unc-service/application/service"
)

type Controllers struct {
	AuthController AuthController
}

func SetupControllers(services *service.Services) *Controllers {
	return &Controllers{
		AuthController: NewAuthController(services.AuthService),
	}
}
