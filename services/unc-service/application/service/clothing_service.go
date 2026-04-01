package service

import (
	"unc/services/unc-service/domain/repository"
)

type ClothingService interface{}

type ClothingServiceImpl struct {
	repository repository.Querier
}

func NewClothingService(repository repository.Querier) ClothingService {
	return &ClothingServiceImpl{
		repository: repository,
	}
}
