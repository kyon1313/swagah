package model

type PreferredFT struct {
	Cid         string `json:"cid"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Insti       string `json:"insti"`
	AccountType string `json:"account_type"`
	DateAdded   string `json:"date_added"`
	TargetCid   string `json:"target_cid"`
	Mobile      string `json:"mobile"`
}

type DeletePreferredAccount struct {
	Cid       string `json:"cid"`
	TargetCid string `json:"target_cid"`
}

type UpdatePreferredFT struct {
	Cid     string `json:"cid"`
	Account string `json:"account"`
}

type BugReport struct {
	ID              uint                   `json:"-" gorm:"primaryKey" `
	InstiCode       int                    `json:"insti_code"`
	ApplicationCode int                    `json:"application_code"`
	Details         map[string]interface{} `json:"details"`
	Cid             int                    `json:"cid"`
	Date            string                 `json:"date"`
	Status          int                    `json:"status"`
}

type GetReports struct {
	Cid    int `json:"cid"`
	Status int `json:"status"`
}

type FeeStructureRequest struct {
	FeeId int `json:"fee_id"`
}

type GetSaveAcc struct {
	Cid string `json:"cid"`
}

type GetSaveAccAndTCid struct {
	Cid       string `json:"cid"`
	TargetCid string `json:"target_cid"`
}

type UpdateBugreports struct {
	Id     uint `json:"id"`
	Status int  `json:"status"`
}

type GetAtmByInsti struct {
	Inst_code string `json:"inst_code"`
}

type SendBugReport struct {
	InstiCode       int                    `json:"insti_code"`
	ApplicationCode int                    `json:"application_code"`
	Details         map[string]interface{} `json:"details"`
	Cid             int                    `json:"cid"`
}
