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
