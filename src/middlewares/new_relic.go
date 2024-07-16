package middlewares

import (
	"os"

	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/fiber/v2"
)

func NewRelic(ctx *fiber.App) fiber.Handler {
	cfg := fibernewrelic.Config{
		License: os.Getenv("NEW_RELIC_LICENSE_KEY"),
		AppName: "go-fiber-newrelic",
		Enabled: true,
	}
	// func(*fiber.Ctx) error
	return fibernewrelic.New(cfg)
}
