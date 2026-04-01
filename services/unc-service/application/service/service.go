package service

import (
	"unc/services/unc-service/application/client"
	"unc/services/unc-service/domain/repository"
)

type Services struct {
	AuthService     AuthService
	ClothingService ClothingService
}

func SetupServices(repository repository.Querier, clients *client.Clients) *Services {
	return &Services{
		AuthService:     NewAuthService(repository, clients.EmailClient),
		ClothingService: NewClothingService(repository),
	}
}
