package router

import (
	"unc/services/unc-service/application/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, controllers *controller.Controllers) {
	api := app.Group("/api/v0")

	setupAuthRoutes(api, controllers.AuthController)
}
