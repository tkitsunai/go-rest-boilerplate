package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/go-rest-boilerplate/internal/config"
	"github.com/tkitsunai/go-rest-boilerplate/internal/rest/routes"
	"net"
)

type App struct {
	engine *fiber.App
	port   string
}

func NewApp() *App {
	app := fiber.New(fiber.Config{
		// Pre-fork mode is thread control of the entire application, not just the thread that handles Fiber, becomes hard.
		// If pre-fork mode is turned true, it needs to be considered throughout the application.
		Prefork: false,
	})
	s := &App{
		engine: app,
		port:   config.Get().Port,
	}
	routes.SetupRoutes(s.engine)
	return s
}

func (s *App) Run() error {
	return s.engine.Listen(net.JoinHostPort("", s.port))
}
