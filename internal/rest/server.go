package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/go-rest-boilerplate/internal/config"
)

type Server struct {
	engine *fiber.App
	port   string
}

func New() *Server {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	s := &Server{
		engine: app,
		port:   config.Get().Port,
	}
	s.registerRoutes()
	return s
}

func (s *Server) Run() error {
	return s.engine.Listen(fmt.Sprintf(":%s", s.port))
}
