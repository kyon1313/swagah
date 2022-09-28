package route

import "github.com/gofiber/fiber/v2"

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
	Principal      int         `json:"principal"`
	Interest       int         `json:"interest"`
	Gives          int         `json:"gives"`
	IbalPrin       int         `json:"ibalPrin"`
	IbalInt        int         `json:"ibalInt"`
	BalPrin        int         `json:"balPrin"`
	BalInt         int         `json:"balInt"`
	Amort          float32     `json:"amort"`
	DuePrin        float32     `json:"duePrin"`
	DueInt         int         `json:"dueInt"`
	LoanBal        int         `json:"loanBal"`
	SaveBal        float32     `json:"saveBal"`
	WaiveInt       int         `json:"waiveInt"`
	UnpaidCtr      int         `json:"unpaidCtr"`
	Writtenoff     int         `json:"writtenoff"`
	Classification int         `json:"classification"`
	ClassDesc      string      `json:"classDesc"`
	Writeoff       int         `json:"writeoff"`
	Pay            float32     `json:"pay"`
	Withdraw       float32     `json:"withdraw"`
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
		return c.SendString("fucking shit")
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
			Principal:      0,
			Interest:       0,
			Gives:          0,
			IbalInt:        0,
			IbalPrin:       0,
			BalPrin:        0,
			BalInt:         0,
			Amort:          50.0,
			DuePrin:        50.0,
			DueInt:         0,
			LoanBal:        0,
			SaveBal:        3772.97,
			WaiveInt:       0,
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
