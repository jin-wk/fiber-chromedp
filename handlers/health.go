package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/jin-wk/fiber-mysql/models"
)

// CheckHealth godoc
// @Summary Check Health for API Server
// @Description Check Health for API Server
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ResponseModel{}
// @Failure 404 {object} models.ResponseModel{}
// @Failure 503 {object} models.ResponseModel{}
// @Router /api/health [get]
func CheckHealth(c *fiber.Ctx) error {
	return c.JSON(models.Response(1000, "success", nil))
}
