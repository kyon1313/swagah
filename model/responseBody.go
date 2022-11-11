package model

import "time"

type GetparamWebtools struct {
	Param_id          int         `json:"param_id"`
	Param_name        string      `json:"param_name"`
	Param_value       string      `json:"param_value"`
	Param_desc        string      `json:"param_desc"`
	App_type          string      `json:"app_type"`
	Created_by        int         `json:"created_by"`
	Created_date      interface{} `json:"created_date"`
	Last_updated_by   interface{} `json:"last_updated_by"`
	Last_updated_date interface{} `json:"last_updated_date"`
	Param_status      string      `json:"param_status"`
	Process_id        interface{} `json:"process_id"`
	Value_type        string      `json:"value_type"`
	Value_lookup      interface{} `json:"value_lookup"`
}

type BugReportResponse struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	InstiCode       int    `json:"insti_code"`
	ApplicationCode int    `json:"application_code"`
	Details         string `json:"details"`
	Cid             int    `json:"cid"`
	Date            string `json:"date"`
	Status          int    `json:"status"`
}

type FeeStructure struct {
	Fee_id              int     `json:"fee_id"`
	Start_range         int     `json:"start_range"`
	End_range           int     `json:"end_range"`
	Total_charge        float64 `json:"total_charge"`
	Agent_income        float64 `json:"agent_income"`
	Fds_fee             float64 `json:"fds_fee"`
	Cmit_fee            float64 `json:"cmit_fee"`
	Bank_income         float64 `json:"bank_income"`
	Created_by          int     `json:"created_by"`
	Created_date        string  `json:"created_date"`
	Last_updated_by     int     `json:"last_updated_by"`
	Last_updated_date   string  `json:"last_updated_date"`
	Trans_type          string  `json:"trans_type"`
	Telco_fee           float64 `json:"telco_fee"`
	Bank_income_flag    bool    `json:"bank_income_flag"`
	Client_type         string  `json:"client_type"`
	Agent_target_income float64 `json:"agent_target_income"`
	Bancnet_income      float64 `json:"bancnet_income"`
}

type BugReportResponses struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	InstiCode       int    `json:"insti_code"`
	ApplicationCode int    `json:"application_code"`
	Cid             int    `json:"cid"`
	Date            string `json:"date"`
	Status          int    `json:"status"`
}

type GetAtm struct {
	Atm_id            uint    `json:"atm_id"`
	Atm_type          string  `json:"atm_type"`
	Atm_address       string  `json:"atm_address"`
	Created_by        uint    `json:"created_by"`
	Created_date      string  `json:"created_date"`
	Last_updated_by   uint    `json:"last_updated_by"`
	Last_updated_date string  `json:"last_updated_date"`
	Atm_city          string  `json:"atm_city"`
	Atm_description   string  `json:"atm_description"`
	Atm_longitude     float64 `json:"atm_longitude"`
	Atm_latitude      float64 `json:"atm_latitude"`
	Inst_code         string  `json:"inst_code"`
}

type Error struct {
	Message string `json:"message"`
	RetCode int    `json:"retcode"`
}

type TimeModel struct {
	TimeServer time.Time `json:"time"`
}
