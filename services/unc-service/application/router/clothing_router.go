package router

import (
	"unc/services/unc-service/application/controller"
	"unc/services/unc-service/application/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupClothingRoutes(api fiber.Router, clothingController controller.ClothingController) {
	clothings := api.Group("/clothings")

	clothings.Post("/create", middleware.AuthMiddleware(), clothingController.Create)
}
