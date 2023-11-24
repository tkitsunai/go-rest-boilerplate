package rest

import "github.com/gofiber/fiber/v2"

func (s *Server) registerRoutes() {
	s.engine.Get("/systems/ping", PingPong)
	s.engine.Get("/systems/stats", Stats)
}

func PingPong(ctx *fiber.Ctx) error {
	return ctx.SendString("pong")
}

func Stats(ctx *fiber.Ctx) error {
	return ctx.SendString("pong")
}
