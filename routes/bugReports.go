package routes

import (
	"encoding/json"
	"fmt"
	"sample/database"
	"sample/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary    Send Bug Report
// @Description  Send But Report
// @Tags         BugReport
// @Accept       json
// @Produce      json
// @Param        user body  model.SendBugReport true "Search"
// @Success      200  {object} model.BugReport
// @Failure      400  {object}  model.Error
// @Router       /SendBugReport [post]
func BugReporting(c *fiber.Ctx) error { //create Bug Report
	var b model.SendBugReport
	if err := c.BodyParser(&b); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	now := time.Now()
	y, m, d := now.Date()

	theTime := fmt.Sprintf("%v %d %d", m, d, y)

	if err := database.DB.Debug().Raw("insert into mfs.json_b (application_code,insti_code,details,cid,date,status) values (?,?,?,?,?,1) ", b.ApplicationCode, b.InstiCode, b.Details, b.Cid, theTime).Scan(&b).Error; err != nil {
		return c.JSON(&fiber.Map{
			"message": "Encountered an error",
			"error":   err,
		})
	}

	return c.JSON(&fiber.Map{
		"data": b,
	})

}

// @Summary    Update Bug Report
// @Description  Update Bug Report
// @Tags         BugReport
// @Accept       json
// @Produce      json
// @Param        user body  model.UpdateBugreports true "Search"
// @Success      200  {object} model.BugReportResponses
// @Failure      400  {object}  model.Error
// @Router       /updateBugreport [post]
func UpdateBugreport(c *fiber.Ctx) error { //update bug report
	var br model.BugReportResponses
	var ubr model.UpdateBugreports

	if err := c.BodyParser(&ubr); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	database.DB.Debug().Table("mfs.json_b").Updates(ubr).Find(&br)

	return c.JSON(&fiber.Map{"data": br})

}

// @Summary    GetBugReports
// @Description  GetBugReports
// @Tags         BugReport
// @Accept       json
// @Produce      json
// @Param        user body  model.GetReports true "Search"
// @Success      200  {object} model.BugReport
// @Failure      400  {object}  model.Error
// @Router       /GetBugReports [post]
func GetBugReports(c *fiber.Ctx) error { //get bug reports
	var gr model.GetReports
	var bugReport []model.BugReportResponse
	var bug []model.BugReport

	if err := c.BodyParser(&gr); err != nil {
		c.SendString("internal server error")
	}

	rows := database.DB.Debug().Table("mfs.json_b").Find(&bugReport, gr).RowsAffected
	if rows == 0 {
		return c.JSON(&fiber.Map{"message": "No data found"})
	}

	for _, v := range bugReport {
		Jsondata := make(map[string]interface{})
		j := json.RawMessage(v.Details)
		err := json.Unmarshal(j, &Jsondata)
		if err != nil {
			return c.JSON(&fiber.Map{
				"message": "Cant unmarshal",
			})
		}
		response := model.GetBugReportConstructor(v, Jsondata)
		bug = append(bug, response)

	}

	return c.JSON(&fiber.Map{"data": bug})
}
