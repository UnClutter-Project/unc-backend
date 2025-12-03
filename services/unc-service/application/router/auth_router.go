package router

import (
	"unc/services/unc-service/application/controller"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(api fiber.Router, authController controller.AuthController) {
	users := api.Group("/users")

	users.Post("/register", authController.Register)
}
