package route

import (
	"bytes"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LoanList struct {
	Acc         string  `json:"acc"`
	Status      string  `json:"status"`
	DataRelease string  `json:"dateRelease"`
	AcctType    string  `json:"acctType"`
	Principal   float32 `json:"principal"`
	Interest    float32 `json:"interest"`
	Oth         float32 `json:"oth"`
	Balance     float32 `json:"balance"`
	Term        int     `json:"term"`
	PaidTerm    int     `json:"paidTerm"`
}

/*

JANUS

/CoreAccounts/API/searchLoanList



*/

// Janus kumware godoc
// @Summary     /CoreAccounts/API/searchLoanList
// @Description  /CoreAccounts/API/searchLoanList
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} LoanList
// @Failure      400  {object}  Errror
// @Router       /janus/loanlist/ [post]
func SearchLoanList(c *fiber.Ctx) error {
	var user Use

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user.Cid)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/searchLoanList", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := []LoanList{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
	}

	return c.JSON(result)

}
