package routes

import (
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

/*
"THIS INCLUDES:
getSpecifiedFee
getAllFees
getInstiFees"
*/

// @Summary    Get All Fees
// @Description  Get All Fees
// @Tags        GetFees
// @Accept       json
// @Produce      json
// @Success      200  {object} model.FeeStructure
// @Failure      400  {object}  model.Error
// @Router       /getAllFees [get]
func GetAllFees(c *fiber.Ctx) error {
	var fs []model.FeeStructure
	database.DB.Raw("Select * from mfs.m_fee_Structure").Scan(&fs)
	return c.JSON(&fiber.Map{"data": fs})

}

// @Summary    GetSpecifiedFee
// @Description  GetSpecifiedFee
// @Tags         GetFees
// @Accept       json
// @Produce      json
// @Param        user body  model.FeeStructureRequest true "Search"
// @Success      200  {object} model.FeeStructure
// @Failure      400  {object}  model.Error
// @Router       /getSpecifiedFee [post]
func GetSpecifiedFee(c *fiber.Ctx) error {
	var fsr model.FeeStructureRequest
	var fs model.FeeStructure
	if err := c.BodyParser(&fsr); err != nil {
		return c.JSON(&fiber.Map{
			"message": "internal server error",
		})
	}
	database.DB.Raw("Select * from mfs.m_fee_Structure Where fee_id=?", fsr.FeeId).Find(&fs)
	return c.JSON(&fiber.Map{"data": fs})
}

// @Summary    GetInstiFees Not done
// @Description  GetInstiFees Not done
// @Tags         GetFees
// @Accept       json
// @Produce      json
// @Param        user body  model.FeeStructureRequest true "Search"
// @Success      200  {object} model.FeeStructure
// @Failure      400  {object}  model.Error
// @Router       / [post]
func GetInstiFees(c *fiber.Ctx) error { return nil }
