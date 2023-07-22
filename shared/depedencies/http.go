package depedencies

import (
	_ "students-tech-server/docs"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
)

func NewHttp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}]:[${ip}] ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.ConfigDefault))

  app.Get("/metrics", monitor.New(monitor.Config{Title: "students tech server metrics"}))

  app.Get("/swagger/*", swagger.HandlerDefault)

	log.Info().Msg("http fiber started")

	return app
}
