package main

import (
	"log"
	"sample/database"
	route "sample/routes"

	_ "sample/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

var match bool

// @title My Fucking Swaggah Jagahh.
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	database.Migration()

	Routes(app)
	app.Get("/docs/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))

}

func Routes(app *fiber.App) {
	proj1 := app.Group("/proj1")

	proj1.Post("/add/", route.Add)
	proj1.Get("/get/:id", route.Get)
	proj1.Get("/gets/", route.GetAll)
	proj1.Delete("/delete/:id", route.Delete)
	proj1.Put("/update/", route.Update)

	///get data from another schema------------------------
	app.Get("/getthis", route.Return2)
	//transaring data from another schema------------------
	app.Get("/updatesss", route.UpdatingData)
	//getting data from another api

	app.Post("/json", route.GEttingMtfks)
	// app.Post("/gettingUsingInterface", route.GetusersUsingDynamicStruct)
	app.Post("/twodata", route.Tryit)
	/*
		janus

	*/
	janus := app.Group("/janus")
	//janus customer savings list
	janus.Post("/customerSavings", route.CustSavings)
	//search loan list using cid
	janus.Post("/loanlist", route.SearchLoanList)
	//getting data from another api with swagger
	janus.Post("/consuming", route.Consuming)
	janus.Post("/searchCID", route.Consuming)

	/*
		swagger hard coded data
	*/
	swag := app.Group("/swag")
	swag.Post("/search/", route.FuckignShit)
	swag.Post("/loanlisthard", route.LoanListHard)
	swag.Post("/custSavingsList", route.CustSavingsHard)
	swag.Post("/custSavingsinfo", route.CustSavingsInfo)
	swag.Post("/loanInfo", route.LoansInfo)
	swag.Post("/generateCol", route.GenerateColCid)
}

func unique(s []int) []int {
	inResult := make(map[int]bool)
	var result []int
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func getUserInfo(userID int) map[string]interface{} {
	user := make(map[string]interface{})
	database.DB.Table("users").Select("user_id, fullname, email, address").Find(&user, "user_id", userID)
	return user
}

func getProductDetails(product_id []uint) []map[string]interface{} {
	// product := make(map[string]interface{})
	result := []map[string]interface{}{}

	// product := []models.Product{}sa
	database.DB.Debug().Table("products p").
		Select("p.product_id, p.name, p.description, p.images, p.stars, c.quantity_ordered").
		Joins("JOIN addtocarts c ON p.product_id = c.product_id").Where("p.product_id IN ?", product_id).Find(&result)
	// fmt.Println(product)
	return result
}
