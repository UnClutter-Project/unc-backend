package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func FormatValidationErrors(err error) []fiber.Map {
	var errors []fiber.Map

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, fiber.Map{
			"field":   strings.ToLower(e.Field()),
			"value":   e.Value(),
			"message": getErrorMessage(e),
		})
	}

	return errors
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "email":
		return err.Field() + " must be a valid email"
	default:
		return err.Field() + " is invalid"
	}
}
