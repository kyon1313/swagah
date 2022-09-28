package model

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
