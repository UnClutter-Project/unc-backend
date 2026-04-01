package router

import (
	"unc/services/unc-service/application/controller"

	"github.com/gofiber/fiber/v2"
)

func setupClothingRoutes(api fiber.Router, clothingController controller.ClothingController) {
	_ = api.Group("/clothing")
}
