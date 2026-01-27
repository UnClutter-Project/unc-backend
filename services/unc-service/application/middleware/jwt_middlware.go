package middleware

import (
	"unc/services/unc-service/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetConfig().JWTSecret)},
		// ErrorHandler: jwtError,
		// SigningKey: []byte(config.GetConfig().JWTSecret),
	})
}

// func jwtError(c fiber.Ctx, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
// 	}
// 	return c.Status(fiber.StatusUnauthorized).
// 		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
// }
