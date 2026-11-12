package handlers

import (
	"github.com/gofiber/fiber/v2"

	"example.com/matrix-processing-system/api-go/internal/models"
	"example.com/matrix-processing-system/api-go/internal/services"
)

// QRHandler procesa la factorización QR y las estadísticas.
func QRHandler(c *fiber.Ctx) error {

	var request models.MatrixRequest

	// Deserializar el JSON recibido.
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid JSON request",
		})
	}

	// Ejecutar la factorización QR.
	q, r, err := services.FactorizeQR(request.Matrix)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// Enviar Q y R a Node.js para calcular las estadísticas.
	statistics, err := services.GetStatistics(q, r)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not calculate matrix statistics",
			"details": err.Error(),
		})
	}

	// Construir la respuesta final.
	response := fiber.Map{
		"q":          q,
		"r":          r,
		"statistics": statistics,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}
