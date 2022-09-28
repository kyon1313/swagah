package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

type JsonStruct struct {
	Cid               int    `json:"cid"`
	Lastname          string `json:"lastname"`
	FirstName         string `json:"firstname"`
	MiddleName        string `json:"middlename"`
	MaidenFName       string `json:"maidenFname"`
	MaidenLName       string `json:"maidenLName"`
	MaidenMName       string `json:"maidenMName"`
	DoBirth           string `json:"doBirth"`
	BirthPlace        string `json:"birthPlace"`
	Sex               int    `json:"sex"`
	CivilStatus       int    `json:"civilStatus"`
	Title             int    `json:"title"`
	Status            int    `json:"status"`
	Classification    int    `json:"classification"`
	SubClassification int    `json:"subClassification"`
	Business          int    `json:"business"`
	DEntry            string `json:"doEntry"`
	DoRecognized      string `json:"doRecognized"`
	DoResigned        string `json:"doResigned"`
	BrCode            int    `json:"brCode"`
	UnitCode          int    `json:"unitCode"`
	CenterCode        int    `json:"centerCode"`
	Dosri             int    `json:"dosri"`
	Reffered          string `json:"reffered"`
	Remarks           string `json:"remarks"`
	AccountNumber     string `json:"accountNumber"`
	SearchName        string `json:"searchName"`

	Address interface{} `json:"address"`
	Contact interface{} `json:"contact"`

	MemberMaidenFName string `json:"memberMaidenFName"`
	MemberMaidenLName string `json:"memberMaidenLName"`
	MemberMaidenMName string `json:"memberMaidenMName"`
	UnitName          string `json:"unitName"`
	CenterName        string `json:"centerName"`
}

type Addres struct {
	AddressID int `json:"addressID"`
	Iiid      int `json:"iiid"`

	AddressTypeID  int    `json:"addressTypeID"`
	AddressType    string `json:"addressType"`
	ProvinceID     int    `json:"provinceID"`
	Province       string `json:"province"`
	TownCity       string `json:"townCity"`
	TowncityID     int    `json:"towncityID"`
	BarangayID     int    `json:"barangayID"`
	Barangay       int    `json:"barangay"`
	ZipCode        int    `json:"zipCode"`
	AddressDetails string `json:"addressDetails"`
}

type Contact struct {
	Iiid          int    `json:"iiid"`
	Series        int    `json:"series"`
	ContactTypeId int    `json:"contactTypeId"`
	ContactType   string `json:"contactType"`
	Contact       string `json:"contact"`
}

type Use struct {
	Cid int `json:"cid" example:"9"`
}

// Janus godoc
// @Summary    SearchCustomerCID
// @Description  SearchCustomerCID
// @Tags         Janus Hard coded data
// @Accept       json
// @Produce      json
// @Param        user body Use true "Search"
// @Success      200  {object} JsonStruct
// @Failure      400  {object}  Errror
// @Router       /swag/search/ [post]
func FuckignShit(c *fiber.Ctx) error {
	var user Use
	u := JsonStruct{
		Cid:               user.Cid,
		Lastname:          "Beligon",
		FirstName:         "Geromme",
		MiddleName:        "coolguy",
		MaidenFName:       "",
		MaidenLName:       "",
		MaidenMName:       "",
		DoBirth:           "1975-08-27",
		BirthPlace:        "Ilo Ilo City",
		Sex:               0,
		CivilStatus:       0,
		Title:             0,
		Status:            0,
		Classification:    0,
		SubClassification: 0,
		Business:          0,
		DEntry:            "      sample2",
		DoRecognized:      "      sample2",
		DoResigned:        "sample 2",
		BrCode:            0,
		UnitCode:          0,
		CenterCode:        0,
		Dosri:             0,
		Reffered:          "      sample2",
		Remarks:           "      sample2",
		AccountNumber:     "      sample2",
		SearchName:        "      sample2",

		Address: []Addres{
			{
				AddressID:      0,
				Iiid:           0,
				Barangay:       0,
				AddressTypeID:  0,
				AddressType:    "      sample2",
				ProvinceID:     0,
				Province:       "      sample2",
				TownCity:       "      sample2",
				TowncityID:     0,
				BarangayID:     0,
				ZipCode:        0,
				AddressDetails: "      sample2",
			},
		},
		Contact: []Contact{
			{
				Iiid:          0,
				Series:        0,
				ContactTypeId: 0,
				ContactType:   "      sample2",
				Contact:       "      sample2",
			},
		},
		MemberMaidenFName: "      sample2",
		MemberMaidenLName: "      sample2",
		MemberMaidenMName: "      sample2",
		UnitName:          "      sample2",
		CenterName:        "      sample2",
	}

	u1 := JsonStruct{
		Cid:               user.Cid,
		Lastname:          "Polintang",
		FirstName:         "Roldan",
		MiddleName:        "Cutieboy",
		MaidenFName:       "",
		MaidenLName:       "",
		MaidenMName:       "",
		DoBirth:           "1975-08-27",
		BirthPlace:        "Ilo Ilo City",
		Sex:               0,
		CivilStatus:       0,
		Title:             0,
		Status:            0,
		Classification:    0,
		SubClassification: 0,
		Business:          0,
		DEntry:            "  sample3",
		DoRecognized:      "  sample3",
		DoResigned:        "sample 3",
		BrCode:            0,
		UnitCode:          0,
		CenterCode:        0,
		Dosri:             0,
		Reffered:          "  sample3",
		Remarks:           "  sample3",
		AccountNumber:     "  sample3",
		SearchName:        "  sample3",

		Address: []Addres{
			{
				AddressID:      0,
				Iiid:           0,
				Barangay:       0,
				AddressTypeID:  0,
				AddressType:    "  sample3",
				ProvinceID:     0,
				Province:       "  sample3",
				TownCity:       "  sample3",
				TowncityID:     0,
				BarangayID:     0,
				ZipCode:        0,
				AddressDetails: "  sample3",
			},
		},
		Contact: []Contact{
			{
				Iiid:          0,
				Series:        0,
				ContactTypeId: 0,
				ContactType:   "  sample3",
				Contact:       "  sample3",
			},
		},
		MemberMaidenFName: "  sample3",
		MemberMaidenLName: "  sample3",
		MemberMaidenMName: "  sample3",
		UnitName:          "  sample3",
		CenterName:        "  sample3",
	}
	u2 := JsonStruct{
		Cid:               user.Cid,
		Lastname:          "Bago",
		FirstName:         "Raymond",
		MiddleName:        "master",
		MaidenFName:       "",
		MaidenLName:       "",
		MaidenMName:       "",
		DoBirth:           "1975-08-27",
		BirthPlace:        "Ilo Ilo City",
		Sex:               0,
		CivilStatus:       0,
		Title:             0,
		Status:            0,
		Classification:    0,
		SubClassification: 0,
		Business:          0,
		DEntry:            "sample",
		DoRecognized:      "sample",
		DoResigned:        "sample ",
		BrCode:            0,
		UnitCode:          0,
		CenterCode:        0,
		Dosri:             0,
		Reffered:          "sample",
		Remarks:           "sample",
		AccountNumber:     "sample",
		SearchName:        "sample",

		Address: []Addres{
			{
				AddressID:      0,
				Iiid:           0,
				Barangay:       0,
				AddressTypeID:  0,
				AddressType:    "sample",
				ProvinceID:     0,
				Province:       "sample",
				TownCity:       "sample",
				TowncityID:     0,
				BarangayID:     0,
				ZipCode:        0,
				AddressDetails: "sample",
			},
		},
		Contact: []Contact{
			{
				Iiid:          0,
				Series:        0,
				ContactTypeId: 0,
				ContactType:   "sample",
				Contact:       "sample",
			},
		},
		MemberMaidenFName: "sample",
		MemberMaidenLName: "sample",
		MemberMaidenMName: "sample",
		UnitName:          "sample",
		CenterName:        "sample",
	}

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("there is an error")
	}
	u.Cid = user.Cid
	u2.Cid = user.Cid
	u1.Cid = user.Cid
	users := []JsonStruct{u, u1, u2}

	random := rand.Intn(len(users))

	return c.JSON(users[random])
}

//----------getting using all fields

func GEttingMtfks(c *fiber.Ctx) error {

	var user model.User
	var search model.User

	if err := c.BodyParser(&search); err != nil {
		return c.SendString("fuck you")
	}
	database.DB.Find(&user, search)
	fmt.Println(user)
	fmt.Println(search)
	return c.JSON(user)
}

/*

	JANUS
/CoreAccounts/API/mobile/api/v1/SearchCustomerCID


*/

//--------------consuming from another api

// Search  godoc
// @Summary     SearchCustomerCID
// @Description  SearchCustomerCID
// @Tags         Janus
// @Accept       json
// @Produce      json
// @Param        user body Use true "SearchbyCid"
// @Success      200  {object} JsonStruct
// @Failure      400  {object}  Errror
// @Router       /janus/searchCID/ [post]
func Consuming(c *fiber.Ctx) error {
	var user JsonStruct

	if err := c.BodyParser(&user); err != nil {
		return c.SendString("fucking shit")
	}

	jsonReq, err := json.Marshal(user.Cid)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/mobile/api/v1/SearchCustomerCID", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return c.SendString("fucking shit")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	result := JsonStruct{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("fucking shit")
	}

	return c.JSON(result)

}

//---------------------------end

type Products struct {
	ProductName string `json:"productName"`
	P           Prices `json:"prices"`
}
type Prices struct {
	Price1 int `json:"price1"`
	Price2 int `json:"price2"`
}

func Tryit(c *fiber.Ctx) error {

	var p Products

	if err := c.BodyParser(&p); err != nil {
		return c.SendString("fucking shit")
	}

	return c.JSON(p)
}
