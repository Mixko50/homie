package fiber

import (
	"github.com/gofiber/fiber/v2"
	"server/router"
	"server/utils/middleware"
)

var app *fiber.App

func Init() {
	// Default config
	app = fiber.New(fiber.Config{
		CaseSensitive: false,
		ErrorHandler:  errorResponse,
		AppName:       "Homie",
	})

	// * Middleware
	app.Use(middleware.Cors())

	// * Router
	app.Route("api/", router.Router)

	// * Not Found
	app.Use(notFound)

	// * Listen
	app.Listen(":8080")
}
