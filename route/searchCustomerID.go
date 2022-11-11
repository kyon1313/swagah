package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

type JsonStruct struct {
	Cid               int         `json:"cid"`
	Lastname          string      `json:"lastname"`
	FirstName         string      `json:"firstname"`
	MiddleName        string      `json:"middlename"`
	MaidenFName       string      `json:"maidenFname"`
	MaidenLName       string      `json:"maidenLName"`
	MaidenMName       string      `json:"maidenMName"`
	DoBirth           string      `json:"doBirth"`
	BirthPlace        string      `json:"birthPlace"`
	Sex               int         `json:"sex"`
	CivilStatus       int         `json:"civilStatus"`
	Title             int         `json:"title"`
	Status            int         `json:"status"`
	Classification    int         `json:"classification"`
	SubClassification int         `json:"subClassification"`
	Business          int         `json:"business"`
	DEntry            string      `json:"doEntry"`
	DoRecognized      string      `json:"doRecognized"`
	DoResigned        string      `json:"doResigned"`
	BrCode            int         `json:"brCode"`
	UnitCode          int         `json:"unitCode"`
	CenterCode        int         `json:"centerCode"`
	Dosri             int         `json:"dosri"`
	Reffered          string      `json:"reffered"`
	Remarks           string      `json:"remarks"`
	AccountNumber     string      `json:"accountNumber"`
	SearchName        interface{} `json:"searchName"`

	Address interface{} `json:"address"`
	Contact interface{} `json:"contact"`

	MemberMaidenFName string      `json:"memberMaidenFName"`
	MemberMaidenLName string      `json:"memberMaidenLName"`
	MemberMaidenMName string      `json:"memberMaidenMName"`
	UnitName          interface{} `json:"unitName"`
	CenterName        interface{} `json:"centerName"`
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
// @Summary     CoreAccounts/API/mobile/api/v1/SearchCustomerCID
// @Description  CoreAccounts/API/mobile/api/v1/SearchCustomerCID
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
		return c.SendString("internal server error")
	}

	jsonReq, err := json.Marshal(user.Cid)

	resp, err := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/mobile/api/v1/SearchCustomerCID", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
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
	result := JsonStruct{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return c.SendString("Request Failed")
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
		return c.SendString("Request Failed")
	}

	return c.JSON(p)
}
