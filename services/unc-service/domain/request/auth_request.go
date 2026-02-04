package request

import (
	"time"
	"unc/services/unc-service/domain/repository"
)

type RegisterRequest struct {
	Username string                `json:"username" validate:"required,min=4,max=32"`
	Email    string                `json:"email" validate:"required,email"`
	Password string                `json:"password" validate:"required,min=8,max=128"`
	Gender   repository.GenderType `json:"gender" validate:"required"`
	DOB      time.Time             `json:"dob" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type VerifyRequest struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Code     string `json:"code" validate:"required,len=6"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
