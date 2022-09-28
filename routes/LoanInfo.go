package route

import "github.com/gofiber/fiber/v2"

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
	Principal           int         `json:"principal"`
	Interest            int         `json:"interest"`
	Others              int         `json:"others"`
	Discounted          int         `json:"discounted"`
	Netproceed          int         `json:"netproceed"`
	Balance             int         `json:"balance"`
	Prin                int         `json:"prin"`
	Intr                int         `json:"intr"`
	Oth                 int         `json:"oth"`
	Penalty             int         `json:"penalty"`
	Waivedint           int         `json:"waivedint"`
	Disbby              string      `json:"disbby"`
	Approvby            string      `json:"approvby"`
	Cycle               int         `json:"cycle"`
	Frequency           int         `json:"frequency"`
	Annumdiv            int         `json:"annumdiv"`
	Lngrpcode           int         `json:"lngrpcode"`
	Proff               int         `json:"proff"`
	Fundsource          int         `json:"fundsource"`
	Conintrate          int         `json:"conintrate"`
	Amortcond           int         `json:"amortcond"`
	Amortcondvalue      int         `json:"amortcondvalue"`
	Classification_code int         `json:"classification_code"`
	Classification_type int         `json:"classification_type"`
	Remarks             interface{} `json:"remarks"`
	Amort               interface{} `json:"amort"`
	IsLumpsum           int         `json:"isLumpsum"`
	LoanID              interface{} `json:"loanID"`
	AmortList           []AmortList `json:"amortList"`
	Charges             Charges     `json:"charges"`
}

type AmortList struct{}
type Charges struct {
	Acc    string    `json:"acc"`
	Charge []Charges `json:"charges"`
}

// Janus godoc
// @Summary    Loan Info
// @Description   Loan Info
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Acc true "Search"
// @Success      200  {object} LaonInfo
// @Failure      400  {object}  Errror
// @Router       /swag/loanInfo [post]
func LoansInfo(c *fiber.Ctx) error {
	var acc Acc
	if err := c.BodyParser(&acc); err != nil {
		return c.SendString("server error")
	}

	a := []LaonInfo{
		{
			Cid:                 0,
			Acc:                 acc.Acc,
			AppType:             0,
			AcctType:            0,
			AccDesc:             "sample",
			Dopen:               "sample",
			Domaturity:          "sample",
			Term:                0,
			Weekspaid:           0,
			Status:              0,
			Principal:           0,
			Interest:            0,
			Others:              0,
			Discounted:          0,
			Netproceed:          0,
			Balance:             0,
			Prin:                0,
			Intr:                0,
			Oth:                 0,
			Penalty:             0,
			Waivedint:           0,
			Disbby:              "sample",
			Approvby:            "sample",
			Cycle:               0,
			Frequency:           0,
			Annumdiv:            0,
			Lngrpcode:           0,
			Proff:               0,
			Fundsource:          0,
			Conintrate:          0,
			Amortcond:           0,
			Amortcondvalue:      0,
			Classification_code: 0,
			Classification_type: 0,
			Remarks:             nil,
			Amort:               nil,
			IsLumpsum:           0,
			LoanID:              nil,
			AmortList:           []AmortList{},
			Charges:             Charges{Acc: acc.Acc, Charge: []Charges{}},
		},
	}
	return c.JSON(a)
}
