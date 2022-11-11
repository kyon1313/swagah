package main

import (
	"log"
	"sample/database"
	"sample/route"
	"sample/routes"

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
		Swag 2nd batch
	*/

	app.Get("/ServerTimestamp", routes.GetServerTimestamp)                                 //done
	app.Get("/ServerDate", routes.GetServerDate)                                           //done
	app.Get("/webtoolParams", routes.GetWbParams)                                          //done
	app.Post("/SendBugReport", routes.BugReporting)                                        //done
	app.Post("/GetBugReports", routes.GetBugReports)                                       //done
	app.Get("/getAllFees", routes.GetAllFees)                                              //done
	app.Post("/getSpecifiedFee", routes.GetSpecifiedFee)                                   //done
	app.Delete("/deleteSavedAccount", routes.DeleteSavedAccount)                           //done
	app.Post("/updateSavedAccount", routes.UpdateSavedAccount)                             //done
	app.Post("/getSavedAccountbyCid", routes.GetSavedAccountbyCid)                         //done
	app.Post("/getSavedAccountbyCidAndtargetCid", routes.GetSavedAccountbyCidAndtargetCid) //done
	app.Post("/updateBugreport", routes.UpdateBugreport)                                   //done
	app.Get("/getAllAtm", routes.GetAllAtm)                                                //done
	app.Post("/getInstiAtm", routes.GetInstiAtm)                                           //done

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
