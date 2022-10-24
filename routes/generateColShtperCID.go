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

// Janus godoc
// @Summary    Generate  Col
// @Description  Generate Col
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} GenerateCol
// @Failure      400  {object}  Errror
// @Router       /swag/generateCol [post]
func GenerateColCid(c *fiber.Ctx) error {
	var use Use
	if err := c.BodyParser(&use); err != nil {
		return c.SendString("Request Failed")
	}

	a := []GenerateCol{
		{AppType: 0,
			Code:           3,
			Status:         3,
			StatusDesc:     "Pastdue: w/ Loans",
			Acc:            "1012-0000-29877093",
			Iiid:           use.Cid,
			Um:             "",
			ClientName:     "boss geromme",
			CenterCode:     9262391,
			CenterName:     "ORO B3",
			ManCode:        915237,
			Unit:           "Sta Ana",
			AreaCode:       467858,
			Area:           "NCR 3",
			StaffName:      "",
			Accttype:       60,
			AcctDesc:       "Pledge Account",
			Disbdate:       "2017-02-09",
			DateStart:      "1900-01-01",
			Maturity:       "1900-01-01",
			Principal:      0.0,
			Interest:       0.0,
			Gives:          0.0,
			IbalInt:        0.0,
			IbalPrin:       0.0,
			BalPrin:        0.0,
			BalInt:         0.0,
			Amort:          50.0,
			DuePrin:        50.0,
			DueInt:         0.0,
			LoanBal:        0.0,
			SaveBal:        3772.97,
			WaiveInt:       0.0,
			UnpaidCtr:      0,
			Writtenoff:     0,
			Classification: 1884,
			ClassDesc:      "Regular Client",
			Writeoff:       0,
			Pay:            0.0,
			Withdraw:       0.0,
			Type:           0,
			Uuid:           nil,
		},
	}
	return c.JSON(a)

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
