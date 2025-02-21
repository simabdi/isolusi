package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func SetupApi() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(cors.New(cors.Config{
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
		}, ","),
	}))

	return app
}
