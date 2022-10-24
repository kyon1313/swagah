package model

import "gorm.io/gorm"

//swagger : model
type User struct {
	//User ID
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	MiddleName string `json:"middlename"`
	Lastname   string `json:"lastname"`
}

type UserParam struct {
	CID int `json:"cid"`
}

// OTP
type Otpcode struct {
	gorm.Model    `json:"-"`
	Id            uint   `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Insti_code    uint   `json:"InstiCode"`
	Mobile        string `json:"mobile" gorm:"index:idx_mobile,unique"`
	Updated_at    uint   `json:"updated_at"`
	Request_count uint   `json:"request_count"`
	Tries_count   uint   `json:"Tries_count"`
	Reset_at      uint   `json:"reset_at"`
	Otp           string `json:"otp"`
	Is_used       bool   `json:"is_used"`
}
