package routes

import (
	"fmt"
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

// @Summary    DeleteSavedAccount
// @Description  DeleteSavedAccount
// @Tags         SavedAccount
// @Accept       json
// @Produce      json
// @Param        user body  model.DeletePreferredAccount true "Search"
// @Success      200  {object} string
// @Failure      400  {object}  model.Error
// @Router       /deleteSavedAccount [delete]
func DeleteSavedAccount(c *fiber.Ctx) error { //delete_preferred_ft
	var dsa model.DeletePreferredAccount
	if err := c.BodyParser(&dsa); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	if dsa.Cid == "" || dsa.TargetCid == "" {
		return c.JSON(&fiber.Map{
			"message": "Provide all fields",
		})
	}

	//check if agent have taht target id

	row := database.DB.Debug().Table("public.c_preferred_ft").Where("cid=? and target_cid=?", dsa.Cid, dsa.TargetCid).Delete(&dsa).RowsAffected
	if row == 0 {
		return c.JSON(&fiber.Map{
			"message": "Target Cid Not Exist",
		})
	}

	fmt.Println(dsa)
	return c.JSON(&fiber.Map{
		"message": "Account successfully deleted",
	})
}

// @Summary    UpdateSavedAccount
// @Description  Fields are optional except cid and target cid
// @Tags         SavedAccount
// @Accept       json
// @Produce      json
// @Param        user body  model.PreferredFT true "Search"
// @Success      200  {object} model.PreferredFT
// @Failure      400  {object}  model.Error
// @Router       /updateSavedAccount [post]
func UpdateSavedAccount(c *fiber.Ctx) error { //create if target cid exist, update if not
	var usa model.PreferredFT
	if err := c.BodyParser(&usa); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	if usa.Cid == "" || usa.TargetCid == "" {
		return c.JSON(&fiber.Map{
			"message": "Provide Cid and target cid",
		})
	}

	rows := database.DB.Debug().Table("public.c_preferred_ft").Where("cid=? and target_cid=?", usa.Cid, usa.TargetCid).Updates(&usa).RowsAffected
	if rows == 0 {
		return c.JSON(&fiber.Map{
			"message": "Update failed check credentials",
		})
	}
	database.DB.Table("public.c_preferred_ft").Where("cid=? and target_cid=?", usa.Cid, usa.TargetCid).Find(&usa)
	return c.JSON(&fiber.Map{
		"data": usa,
	})

}

// @Summary    GetSavedAccountbyCid
// @Description  GetSavedAccountbyCid
// @Tags         SavedAccount
// @Accept       json
// @Produce      json
// @Param        user body  model.GetSaveAcc true "Search"
// @Success      200  {object} model.PreferredFT
// @Failure      400  {object}  model.Error
// @Router       /getSavedAccountbyCid [post]
func GetSavedAccountbyCid(c *fiber.Ctx) error {
	var gsa model.GetSaveAcc
	var sa []model.PreferredFT
	if err := c.BodyParser(&gsa); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	database.DB.Table("public.c_preferred_ft").Where("cid=?", gsa.Cid).Find(&sa)
	return c.JSON(&fiber.Map{
		"data": sa,
	})
}

// @Summary    GetSavedAccount by Cid and target Cid
// @Description  GetSavedAccount by Cid and target Cid
// @Tags         SavedAccount
// @Accept       json
// @Produce      json
// @Param        user body  model.GetSaveAccAndTCid true "Search"
// @Success      200  {object} model.PreferredFT
// @Failure      400  {object}  model.Error
// @Router       /getSavedAccountbyCidAndtargetCid [post]
func GetSavedAccountbyCidAndtargetCid(c *fiber.Ctx) error {
	var gsa model.GetSaveAccAndTCid
	var sa model.PreferredFT
	if err := c.BodyParser(&gsa); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	if gsa.Cid == "" || gsa.TargetCid == "" {
		return c.JSON(&fiber.Map{
			"message": "Provide Cid and target cid",
		})
	}
	database.DB.Table("public.c_preferred_ft").Where("cid=? and target_cid=?", gsa.Cid, gsa.TargetCid).Find(&sa)
	return c.JSON(&fiber.Map{
		"data": sa,
	})
}
