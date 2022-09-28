package route

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type CustSavingInfo struct {
	Cid                int     `json:"cid"`
	Acc                string  `json:"acc"`
	AppType            int     `json:"appType"`
	AcctType           int     `json:"acctType"`
	AccDesc            string  `json:"accDesc"`
	Dopen              string  `json:"dopen"`
	Status             int     `json:"status"`
	ClassificationCode int     `json:"classificationCode"`
	ClassificationType int     `json:"classificationType"`
	Balance            float64 `json:"balance"`
}

type Acc struct {
	Acc string `json:"acc"`
}

// Janus godoc
// @Summary    Customer Savings Info
// @Description  Customer Savings Info
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Acc true "Search"
// @Success      200  {object} CustSavingInfo
// @Failure      400  {object}  Errror
// @Router       /swag/custSavingsinfo [post]
func CustSavingsInfo(c *fiber.Ctx) error {
	var acc Acc
	if err := c.BodyParser(&acc); err != nil {
		return c.SendString("server error")
	}
	a1 := []CustSavingInfo{
		{
			Cid:                0,
			Acc:                acc.Acc,
			AppType:            0,
			AcctType:           0,
			AccDesc:            "sample",
			Dopen:              "sample",
			Status:             0,
			ClassificationCode: 0,
			ClassificationType: 0,
			Balance:            0.0,
		},
	}
	a2 := []CustSavingInfo{
		{
			Cid:                0,
			Acc:                acc.Acc,
			AppType:            0,
			AcctType:           0,
			AccDesc:            "sample2",
			Dopen:              "sample2",
			Status:             0,
			ClassificationCode: 0,
			ClassificationType: 0,
			Balance:            0.0,
		},
	}
	a3 := []CustSavingInfo{
		{
			Cid:                0,
			Acc:                acc.Acc,
			AppType:            0,
			AcctType:           0,
			AccDesc:            "sample3",
			Dopen:              "sample3",
			Status:             0,
			ClassificationCode: 0,
			ClassificationType: 0,
			Balance:            0.0,
		},
	}

	r := [][]CustSavingInfo{a1, a2, a3}

	random := rand.Intn(len(r))
	return c.JSON(r[random])
}
