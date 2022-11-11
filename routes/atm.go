package routes

import (
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

/*
getAllAtm
getInstiAtm
*/

//i change the null from orig janus to string (null=="")

// @Summary    Get Atm
// @Description  Get Atm
// @Tags         GetAtm
// @Accept       json
// @Produce      json
// @Success      200  {object} model.GetAtm
// @Failure      400  {object}  model.Error
// @Router       /getAllAtm [get]
func GetAllAtm(c *fiber.Ctx) error {
	var gaa []model.GetAtm
	database.DB.Table("mfs.m_atm").Find(&gaa)

	return c.JSON(gaa)
}

// @Summary    Get Insti Atm
// @Description  Get Insti Atm
// @Tags         GetAtm
// @Accept       json
// @Produce      json
// @Param        user body  model.GetAtmByInsti true "Search"
// @Success      200  {object} model.GetAtm
// @Failure      400  {object}  model.Error
// @Router       /getInstiAtm [post]
func GetInstiAtm(c *fiber.Ctx) error {
	var gAtm model.GetAtmByInsti
	var gaa []model.GetAtm
	if err := c.BodyParser(&gAtm); err != nil {
		return c.JSON(model.Error{Message: "Internal server error", RetCode: 500})
	}

	database.DB.Table("mfs.m_atm").Find(&gaa, "inst_code=?", gAtm.Inst_code)
	if len(gaa) == 0 {
		return c.JSON(&fiber.Map{
			"message": "no data found",
		})
	}
	return c.JSON(&fiber.Map{
		"data": gaa,
	})
}
