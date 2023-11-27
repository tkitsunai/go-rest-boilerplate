package handlers

import (
	"fmt"
	stats "github.com/fukata/golang-stats-api-handler"
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/go-rest-boilerplate/internal/logger"
)

func PingPong(ctx *fiber.Ctx) error {
	logger.Get().Debug().Msg("pong")
	return ctx.SendString("pong")
}

func Stats(ctx *fiber.Ctx) error {
	st := stats.GetStats()
	logger.Get().Debug().Msg(fmt.Sprintf("%+v", st))
	return ctx.JSON(st)
}
