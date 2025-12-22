package service

import (
	"context"
	"errors"
	"unc/services/unc-service/application/helper"
	"unc/services/unc-service/domain/repository"
	"unc/services/unc-service/domain/request"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService interface {
	Register(ctx context.Context, request *request.RegisterRequest) error
	Login(ctx context.Context, request *request.LoginRequest) (string, error)
}

type AuthServiceImpl struct {
	repository repository.Querier
}

func NewAuthService(repository repository.Querier) AuthService {
	return &AuthServiceImpl{
		repository: repository,
	}
}

func (s *AuthServiceImpl) Register(ctx context.Context, request *request.RegisterRequest) error {
	_, err := s.repository.GetUserByUsernameAndEmail(ctx, &repository.GetUserByUsernameAndEmailParams{
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	_, err = s.repository.CreateUser(ctx, &repository.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		Gender:   request.Gender,
		Dob:      pgtype.Date{Time: request.DOB, Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, request *request.LoginRequest) (string, error) {
	user, err := s.repository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return "", err
	}

	err = helper.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := helper.GenerateToken(user.ID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}
