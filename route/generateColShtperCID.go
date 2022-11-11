package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GenerateCol struct {
	AppType        int         `json:"appType"`
	Code           int         `json:"code"`
	Status         int         `json:"status"`
	StatusDesc     string      `json:"statusDesc"`
	Acc            string      `json:"acc"`
	Iiid           int         `json:"iiid"`
	Um             string      `json:"um"`
	ClientName     string      `json:"clientName"`
	CenterCode     int         `json:"centerCode"`
	CenterName     string      `json:"centerName"`
	ManCode        int         `json:"manCode"`
	Unit           string      `json:"unit"`
	AreaCode       int         `json:"areaCode"`
	Area           string      `json:"area"`
	StaffName      string      `json:"staffName"`
	Accttype       int         `json:"accttype"`
	AcctDesc       string      `json:"acctDesc"`
	Disbdate       string      `json:"disbdate"`
	DateStart      string      `json:"dateStart"`
	Maturity       string      `json:"maturity"`
	Principal      float64     `json:"principal"`
	Interest       float64     `json:"interest"`
	Gives          int         `json:"gives"`
	IbalPrin       float64     `json:"ibalPrin"`
	IbalInt        float64     `json:"ibalInt"`
	BalPrin        float64     `json:"balPrin"`
	BalInt         float64     `json:"balInt"`
	Amort          float64     `json:"amort"`
	DuePrin        float64     `json:"duePrin"`
	DueInt         float64     `json:"dueInt"`
	LoanBal        float64     `json:"loanBal"`
	SaveBal        float64     `json:"saveBal"`
	WaiveInt       float64     `json:"waiveInt"`
	UnpaidCtr      int         `json:"unpaidCtr"`
	Writtenoff     int         `json:"writtenoff"`
	Classification int         `json:"classification"`
	ClassDesc      string      `json:"classDesc"`
	Writeoff       int         `json:"writeoff"`
	Pay            float64     `json:"pay"`
	Withdraw       float64     `json:"withdraw"`
	Type           int         `json:"type"`
	Uuid           interface{} `json:"uuid"`
}

// Janus kumware godoc
// @Summary     CoreAccounts/API/generateColShtperCID
// @Description  CoreAccounts/API/generateColShtperCID
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} GenerateCol
// @Failure      400  {object}  Errror
// @Router       /janus/generateCol/ [post]
func GenerateColCidJanus(c *fiber.Ctx) error {

	var user Use

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/generateColShtperCID", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := []GenerateCol{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
	}

	return c.JSON(result)
}
