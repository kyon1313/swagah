package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary    Get Server Date
// @Description  Get Server Date
// @Tags         ServerTime
// @Accept       json
// @Produce      json
// @Success      200  {object} time.Time
// @Failure      400  {object}  model.Error
// @Router       /ServerDate [get]
func GetServerDate(c *fiber.Ctx) error {
	now := time.Now()
	y, m, d := now.Date()
	theTime := fmt.Sprintf("%v %d %d", m, d, y)

	return c.JSON(&fiber.Map{
		"time": theTime,
	})
}
