package route

import (
	"bytes"
	"fmt"

	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
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

// Janus godoc
// @Summary    LoanList
// @Description  Loanlist
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} LoanList
// @Failure      400  {object}  Errror
// @Router       /swag/loanlisthard [post]
func LoanListHard(c *fiber.Ctx) error {
	var use Use
	fmt.Println(use)

	ll := []LoanList{
		{
			Acc:         "geromme acc",
			Status:      "sample",
			DataRelease: "sample",
			AcctType:    "sample",
			Principal:   0,
			Interest:    0,
			Oth:         0,
			Balance:     0.0,
			Term:        0,
			PaidTerm:    0,
		},
	}
	l2 := []LoanList{
		{
			Acc:         "dan acc",
			Status:      "sample2",
			DataRelease: "sample2",
			AcctType:    "sample2",
			Principal:   0,
			Interest:    0,
			Oth:         0,
			Balance:     0.0,
			Term:        0,
			PaidTerm:    0,
		},
	}
	l3 := []LoanList{
		{
			Acc:         "mond acc",
			Status:      "sample3",
			DataRelease: "sample3",
			AcctType:    "sample3",
			Principal:   0,
			Interest:    0,
			Oth:         0,
			Balance:     0.0,
			Term:        0,
			PaidTerm:    0,
		},
	}

	r := [][]LoanList{ll, l2, l3}
	random := rand.Intn(len(r))

	return c.JSON(r[random])
}
