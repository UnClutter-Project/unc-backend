package setup

import (
	"context"
	"errors"
	"fmt"
	"log"
	"unc/services/unc-service/application/controller"
	"unc/services/unc-service/application/router"
	"unc/services/unc-service/application/service"
	"unc/services/unc-service/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	app  *fiber.App
	port string
}

func NewApp() *App {
	ctx := context.Background()
	app := fiber.New(fiber.Config{
		AppName:      "unc",
		ErrorHandler: customErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	db := setupDB(ctx)
	repository := setupRepository(db)
	services := service.SetupServices(repository)
	controllers := controller.SetupControllers(services, validator.New())
	router.SetupRoutes(app, controllers)

	return &App{
		app:  app,
		port: config.GetConfig().AppPort,
	}
}

func (a *App) Start() error {
	log.Printf("Server starting on port %s", a.port)
	return a.app.Listen(fmt.Sprintf(":%s", a.port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
