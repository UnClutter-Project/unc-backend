package router

import (
	"unc/services/unc-service/application/controller"
	"unc/services/unc-service/application/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(api fiber.Router, authController controller.AuthController) {
	users := api.Group("/users")

	users.Post("/register", authController.Register)
	users.Post("/login", authController.Login)
	users.Post("/verify", authController.Verify)
	users.Post("/refresh", authController.Refresh)
	users.Get("/logout", middleware.AuthMiddleware(), authController.Logout)
}
