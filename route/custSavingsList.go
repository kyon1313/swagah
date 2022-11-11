package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CustSavingsListss struct {
	Cid        int     `json:"cid"`
	Acc        string  `json:"acc"`
	Acctype    int     `json:"acctType"`
	AccDesc    string  `json:"accDesc"`
	Dopen      string  `json:"dopen"`
	StatusDesc string  `json:"statusDesc"`
	Balance    float32 `json:"balance"`
	Status     int     `json:"status"`
}

/*

Janus
/CoreAccounts/API/custSavingsList

*/

// Janus kumware godoc
// @Summary     /CoreAccounts/API/custSavingsList
// @Description  /CoreAccounts/API/custSavingsList
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} CustSavingsListss
// @Failure      400  {object}  Errror
// @Router       /janus/customerSavings/ [post]
func CustSavings(c *fiber.Ctx) error {
	var user Use

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user.Cid)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingsList", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := []CustSavingsListss{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
	}
	return c.JSON(result)

}
