package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CustSavingInfo struct {
	Cid                int         `json:"cid"`
	Acc                interface{} `json:"acc"`
	AppType            int         `json:"appType"`
	AcctType           int         `json:"acctType"`
	AccDesc            string      `json:"accDesc"`
	Dopen              string      `json:"dopen"`
	Status             int         `json:"status"`
	ClassificationCode int         `json:"classificationCode"`
	ClassificationType int         `json:"classificationType"`
	Balance            float64     `json:"balance"`
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

// Janus kumware godoc
// @Summary     Loan info
// @Description  Loan info
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Acc true "Search"
// @Success      200  {object} CustSavingInfo
// @Failure      400  {object}  Errror
// @Router       /janus/custSavingInfoJanus/ [post]
func CustSavingInfoJanus(c *fiber.Ctx) error {
	var user Acc

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("fucking shit")
	}

	jsonReq, err := json.Marshal(user)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingInfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return c.SendString("fucking shit")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	result := []CustSavingInfo{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("fucking shit")
	}

	return c.JSON(result)
}
