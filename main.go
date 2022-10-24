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

	janus := app.Group("/janus")
	//janus customer savings list
	janus.Post("/customerSavings", route.CustSavings)
	//search loan list using cid
	janus.Post("/loanlist", route.SearchLoanList)
	///CoreAccounts/API/mobile/api/v1/SearchCustomerCID
	janus.Post("/consuming", route.Consuming)
	janus.Post("/searchCID", route.Consuming)
	janus.Post("/loanInfoJanus", route.LoanInfoJanus)
	janus.Post("/custSavingInfoJanus", route.CustSavingInfoJanus)
	janus.Post("/generateCol", route.GenerateColCidJanus)
	janus.Post("/OpenPaymentTransaction", route.OpenPayment)
	janus.Post("/multiplePayment", route.MultiplePayment)
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
