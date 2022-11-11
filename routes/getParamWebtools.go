package routes

import (
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

// @Summary    Get WebTool Parameters
// @Description  Get WebTool Parameters
// @Tags         GetWebToolParams
// @Accept       json
// @Produce      json
// @Success      200  {object} model.GetparamWebtools
// @Failure      400  {object}  model.Error
// @Router       /webtoolParams [get]
func GetWbParams(c *fiber.Ctx) error {
	var wt []model.GetparamWebtools
	err := database.DB.Debug().Table("mfs.m_param_config").Find(&wt).Error
	if err != nil {
		return c.SendString("There is an errors")
	}
	return c.JSON(wt)
}
