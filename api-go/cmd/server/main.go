package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"example.com/matrix-processing-system/api-go/internal/handlers"
)

func main() {

	app := fiber.New()

	app.Use(recover.New()) // evita que un panic tumbe el servidor
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Go API is running",
			"service": "matrix-processing-api",
		})
	})

	app.Post("/api/v1/qr", handlers.QRHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Fatal(app.Listen(":" + port))
}
