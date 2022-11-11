package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OpenPaymentRequest struct {
	AccountNumber       string `json:"accountNumber"`
	Amount              int    `json:"amount"`
	TrnReference        string `json:"trnReference"`
	Particulars         string `json:"particulars"`
	Username            string `json:"username"`
	SourceSaveAcc       string `json:"sourceSaveAcc"`
	TransFee            int    `json:"transFee"`
	TransFeeParticulars string `json:"transFeeParticulars"`
}

type OpenPaymentResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Details []Details `json:"details"`
}
type Details struct {
	SourceId            interface{} `json:"sourceId"`
	SourceAccountNumber interface{} `json:"sourceAccountNumber"`
	CustomerNumber      string      `json:"customerNumber"`
	CustomerName        string      `json:"customerName"`
	TransactionAmount   float32     `json:"transactionAmount"`
	ReferenceNumber     string      `json:"referenceNumber"`
	CoreReference       string      `json:"coreReference"`
	PaidDate            interface{} `json:"paidDate"`
}

// Janus kumware godoc
// @Summary     /CoreAccounts/API/OpenPaymentTransaction
// @Description  /CoreAccounts/API/OpenPaymentTransaction
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body OpenPaymentRequest true "Search"
// @Success      200  {object} OpenPaymentResponse
// @Failure      400  {object}  Errror
// @Router       /janus/OpenPaymentTransaction/ [post]
func OpenPayment(c *fiber.Ctx) error {

	var user OpenPaymentRequest

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/OpenPaymentTransaction", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := OpenPaymentResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
	}

	return c.JSON(result)
}
