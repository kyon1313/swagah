package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MultiplePaymentRequest struct {
	Paymant         interface{} `json:"payment"`
	Trndate         string      `json:"trndate"`
	OrNumber        int         `json:"orNumber"`
	PrNumber        string      `json:"prNumber"`
	SourceId        int         `json:"sourceId"`
	Username        string      `json:"username"`
	Particulars     string      `json:"particulars"`
	RemitterCID     string      `json:"remitterCID"`
	TotalCollection int         `json:"totalCollection"`
}

type Payment struct {
	Um             string  `json:"um"`
	Acc            string  `json:"acc"`
	Pay            int     `json:"pay"`
	Area           string  `json:"area"`
	Code           int     `json:"code"`
	Iiid           int     `json:"iiid"`
	Type           string  `json:"type"`
	Unit           string  `json:"unit"`
	Uuid           string  `json:"uuid"`
	Amort          int     `json:"amort"`
	Gives          int     `json:"gives"`
	BalInt         float32 `json:"balInt"`
	DueInt         int     `json:"dueInt"`
	Status         int     `json:"status"`
	AppType        int     `json:"appType"`
	BalPrin        float32 `json:"balPrin"`
	DuePrin        int     `json:"duePrin"`
	IbalInt        float32 `json:"ibalInt"`
	LoanBal        int     `json:"loanBal"`
	ManCode        int     `json:"manCode"`
	SaveBal        int     `json:"saveBal"`
	AcctDesc       string  `json:"acctDesc"`
	Accttype       int     `json:"accttype"`
	AreaCode       int     `json:"areaCode"`
	Disbdate       string  `json:"disbdate"`
	IbalPrin       float32 `json:"ibalPrin"`
	Interest       int     `json:"interest"`
	Maturity       string  `json:"maturity"`
	WaiveInt       int     `json:"waiveInt"`
	Withdraw       int     `json:"withdraw"`
	Writeoff       int     `json:"writeoff"`
	ClassDesc      string  `json:"classDesc"`
	DateStart      string  `json:"dateStart"`
	Principal      int     `json:"principal"`
	StaffName      string  `json:"staffName"`
	UnpaidCtr      int     `json:"unpaidCtr"`
	CenterCode     int     `json:"centerCode"`
	CenterName     string  `json:"centerName"`
	ClientName     string  `json:"clientName"`
	StatusDesc     string  `json:"statusDesc"`
	Writtenoff     int     `json:"writtenoff"`
	Classification int     `json:"classification"`
}

type MultiplePaymentResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
}

// Janus kumware godoc
// @Summary     /CoreMFS/API/multiplePayment
// @Description  /CoreMFS/API/multiplePayment
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body MultiplePaymentRequest true "Search"
// @Success      200  {object} MultiplePaymentResponse
// @Failure      400  {object}  Errror
// @Router       /janus/multiplePayment/ [post]
func MultiplePayment(c *fiber.Ctx) error {

	var user MultiplePaymentRequest

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error ")
	}
	fmt.Println(user)
	jsonReq, err := json.Marshal(user)
	resp, err := http.Post("https://cmfstest.cardmri.com/CoreMFS/API/multiplePayment", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return c.SendString("request failed")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	result := MultiplePaymentResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("request failed")
	}

	return c.JSON(result)
}
