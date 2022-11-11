package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LaonInfo struct {
	Cid                 int         `json:"cid"`
	Acc                 string      `json:"acc"`
	AppType             int         `json:"appType"`
	AcctType            int         `json:"acctType"`
	AccDesc             string      `json:"accDesc"`
	Dopen               string      `json:"dopen"`
	Domaturity          string      `json:"domaturity"`
	Term                int         `json:"term"`
	Weekspaid           int         `json:"weekspaid"`
	Status              int         `json:"status"`
	Principal           float64     `json:"principal"`
	Interest            float64     `json:"interest"`
	Others              float64     `json:"others"`
	Discounted          float64     `json:"discounted"`
	Netproceed          float64     `json:"netproceed"`
	Balance             float64     `json:"balance"`
	Prin                float64     `json:"prin"`
	Intr                float64     `json:"intr"`
	Oth                 float64     `json:"oth"`
	Penalty             float64     `json:"penalty"`
	Waivedint           float64     `json:"waivedint"`
	Disbby              string      `json:"disbby"`
	Approvby            string      `json:"approvby"`
	Cycle               int         `json:"cycle"`
	Frequency           int         `json:"frequency"`
	Annumdiv            int         `json:"annumdiv"`
	Lngrpcode           int         `json:"lngrpcode"`
	Proff               int         `json:"proff"`
	Fundsource          int         `json:"fundsource"`
	Conintrate          float64     `json:"conintrate"`
	Amortcond           int         `json:"amortcond"`
	Amortcondvalue      float64     `json:"amortcondvalue"`
	Classification_code int         `json:"classification_code"`
	Classification_type int         `json:"classification_type"`
	Remarks             interface{} `json:"remarks"`
	Amort               interface{} `json:"amort"`
	IsLumpsum           int         `json:"isLumpsum"`
	LoanID              interface{} `json:"loanID"`
	AmortList           []AmortList `json:"amortList"`
	Charges             interface{} `json:"charges"`
}

type AmortList struct {
	Dnum        int     `json:"dnum"`
	Acc         string  `json:"acc"`
	DueDate     string  `json:"dueDate"`
	InstFlag    int     `json:"instFlag"`
	Prin        float64 `json:"prin"`
	Intr        float64 `json:"intr"`
	Oth         float64 `json:"oth"`
	Penalty     float64 `json:"penalty"`
	EndBal      float64 `json:"endBal"`
	EndInt      float64 `json:"endInt"`
	EndOth      float64 `json:"endOth"`
	InstPd      float64 `json:"instPd"`
	PenPd       float64 `json:"penPd"`
	CarVal      float64 `json:"carVal"`
	UpInt       float64 `json:"upInt"`
	ServFee     float64 `json:"servFee"`
	PledgeAmort float64 `json:"pledgeAmort"`
}
type Charges struct {
	Acc    interface{} `json:"acc"`
	Charge []Charges   `json:"charges"`
}

// Janus kumware godoc
// @Summary     /CoreAccounts/API/LoanInfo
// @Description  /CoreAccounts/API/LoanInfo
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Acc true "Search"
// @Success      200  {object} LaonInfo
// @Failure      400  {object}  Errror
// @Router       /janus/loanInfoJanus/ [post]
func LoanInfoJanus(c *fiber.Ctx) error {
	var user Acc

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/LoanInfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := []LaonInfo{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
	}

	return c.JSON(result)
}
