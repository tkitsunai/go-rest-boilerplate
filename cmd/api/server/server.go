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
		Prefork: true,
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
