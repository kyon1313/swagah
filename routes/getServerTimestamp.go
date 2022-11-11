package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary    Get Server Timestamp
// @Description  Get Server Timestamp
// @Tags         ServerTime
// @Accept       json
// @Produce      json
// @Success      200  {object} time.Time
// @Failure      400  {object}  model.Error
// @Router       /ServerTimestamp [get]
func GetServerTimestamp(c *fiber.Ctx) error {
	t := time.Now()

	return c.JSON(&fiber.Map{
		"time": t,
	})
}
