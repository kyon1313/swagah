package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Errror struct {
	ErrorMessage string `json:"errormessage"  `
}

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

// Janus kumware godoc
// @Summary     /CoreAccounts/API/custSavingInfo
// @Description  /CoreAccounts/API/custSavingInfo
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
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingInfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return c.SendString("Request Failed")
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
		return c.SendString("Request Failed")
	}

	return c.JSON(result)
}
