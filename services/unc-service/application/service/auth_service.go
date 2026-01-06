package service

import (
	"context"
	"errors"
	"fmt"
	"unc/services/unc-service/application/client"
	"unc/services/unc-service/application/helper"
	"unc/services/unc-service/domain/repository"
	"unc/services/unc-service/domain/request"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService interface {
	Register(ctx context.Context, request *request.RegisterRequest) error
	Login(ctx context.Context, request *request.LoginRequest) (string, error)
	Verify(ctx context.Context, request *request.VerifyRequest) error
}

type AuthServiceImpl struct {
	repository  repository.Querier
	emailClient client.EmailClient
}

func NewAuthService(repository repository.Querier, emailClient client.EmailClient) AuthService {
	return &AuthServiceImpl{
		repository:  repository,
		emailClient: emailClient,
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

	user, err := s.repository.CreateUser(ctx, &repository.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		Gender:   request.Gender,
		Dob:      pgtype.Date{Time: request.DOB, Valid: true},
	})
	if err != nil {
		return err
	}

	err = s.sendOTP(ctx, user)
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

func (s *AuthServiceImpl) Verify(ctx context.Context, request *request.VerifyRequest) error {
	user, err := s.repository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return err
	}

	_, err = s.repository.GetValidOTPByUsernameAndCode(ctx, &repository.GetValidOTPByUsernameAndCodeParams{
		UserID: user.ID,
		Code:   request.Code,
	})
	if err != nil {
		return err
	}

	_, err = s.repository.SetIsVerifiedByUsername(ctx, &repository.SetIsVerifiedByUsernameParams{
		IsVerified: true,
		Username:   request.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) sendOTP(ctx context.Context, user *repository.Users) error {
	code, err := helper.GenerateOTPCode()
	if err != nil {
		return err
	}

	_, err = s.repository.CreateOTP(ctx, &repository.CreateOTPParams{
		UserID: user.ID,
		Code:   code,
	})

	err = s.emailClient.SendOTP(ctx, user.Email, user.Email, fmt.Sprintf("Take it or leave it: %s", code))
	if err != nil {
		return err
	}

	return nil
}
