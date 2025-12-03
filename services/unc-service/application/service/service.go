package service

import "unc/services/unc-service/domain/repository"

type Services struct {
	AuthService AuthService
}

func SetupServices(repository repository.Querier) *Services {
	return &Services{
		AuthService: NewAuthService(repository),
	}
}
