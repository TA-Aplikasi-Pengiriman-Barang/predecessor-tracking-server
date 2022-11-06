package depedencies

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewHttp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}]:[${ip} - ${requestid}] ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Bikun Tracking Metrics"}))

	return app
}
