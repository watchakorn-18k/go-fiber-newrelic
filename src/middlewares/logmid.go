package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLogger(ctx *fiber.App) fiber.Handler {
	cfg := logger.Config{
		TimeFormat: "15:04:05",
		TimeZone:   "Asia/Bangkok",
	}
	return logger.New(cfg)
}
