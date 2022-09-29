package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
// @Summary     Customer Savings List
// @Description  Customer Savings List
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
		return c.SendString("internal shit")
	}

	jsonReq, err := json.Marshal(user.Cid)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingsList", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := []CustSavingsListss{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("fucking shit")
	}
	return c.JSON(result)

}

// Janus godoc
// @Summary    Customer Savings List
// @Description  Customer Savings List
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} CustSavingsListss
// @Failure      400  {object}  Errror
// @Router       /swag/custSavingsList [post]
func CustSavingsHard(c *fiber.Ctx) error {
	var use Use
	fmt.Println(use)
	if err := c.BodyParser(&use); err != nil {
		return c.SendString("fucking shit")
	}
	cust := []CustSavingsListss{
		{
			Cid:        use.Cid,
			Acc:        "geromme acc",
			Acctype:    0,
			AccDesc:    "sample",
			Dopen:      "sample",
			StatusDesc: "sample",
			Balance:    0.0,
			Status:     0,
		},
	}
	cust2 := []CustSavingsListss{
		{
			Cid:        use.Cid,
			Acc:        "dan acc",
			Acctype:    0,
			AccDesc:    "sample",
			Dopen:      "sample",
			StatusDesc: "sample",
			Balance:    0.0,
			Status:     0,
		},
	}
	cust3 := []CustSavingsListss{
		{
			Cid:        use.Cid,
			Acc:        "mond acc",
			Acctype:    0,
			AccDesc:    "sample",
			Dopen:      "sample",
			StatusDesc: "sample",
			Balance:    0.0,
			Status:     0,
		},
	}
	customers := [][]CustSavingsListss{cust, cust2, cust3}
	random := rand.Intn(len(customers))
	fmt.Print()
	return c.JSON(customers[random])

}
