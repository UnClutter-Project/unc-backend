package service

import "unc/services/unc-service/domain/repository"

type AuthService interface {
	Register(username, email, password string) error
}

type AuthServiceImpl struct {
	repository repository.Querier
}

func NewAuthService(repository repository.Querier) AuthService {
	return &AuthServiceImpl{
		repository: repository,
	}
}

func (s *AuthServiceImpl) Register(username, email, password string) error {
	return nil
}
