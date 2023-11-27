package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/go-rest-boilerplate/internal/rest/handlers"
)

func SetupRoutes(app *fiber.App) {
	setupRoutes(app)
}

func setupRoutes(app *fiber.App) {
	app.Get("/systems/ping", handlers.PingPong)
	app.Get("/systems/stats", handlers.Stats)
}
