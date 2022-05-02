package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// CheckHealth godoc
// @Summary Check Health for API Server
// @Description Check Health for API Server
// @Tags health
// @Accept  json
// @Produce  json
// @Router /api/health [get]
func CheckHealth(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "Ok",
		"data":    nil,
	})
}
