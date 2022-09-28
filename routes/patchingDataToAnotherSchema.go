package route

import (
	"fmt"
	"sample/database"

	"github.com/gofiber/fiber/v2"
)

type Users struct {
	Name string `json:"name"`
}
type User2 struct {
	Name string `json:"name"`
}

//transfer data from another schema
func UpdatingData(c *fiber.Ctx) error {
	var user []Users
	var mf2user []User2
	var newUsers []Users

	//public.users
	database.DB.Debug().Raw("Select * from users").Scan(&user)
	fmt.Println(user)
	//mfs.user
	database.DB.Debug().Raw("Select * from mfs.user").Scan(&mf2user)
	fmt.Println(mf2user)

	count := 0
	if len(user) == 0 {
		for _, v := range mf2user {
			user = append(user, Users{v.Name})
			count++
		}
		database.DB.Debug().Create(&user)
	} else {
		for i := 0; i < len(mf2user); {
			exist := true
			for _, b := range user {
				if b.Name == mf2user[i].Name {
					exist = false
					break
				}
			}
			if !exist {
				mf2user = append(mf2user[:i], mf2user[i+1:]...) // delete names[i]

			} else {
				i++
			}

		}
		fmt.Println("not on the list users", mf2user)

		for _, v := range mf2user {
			newUsers = append(newUsers, Users{v.Name})
		}
		if len(newUsers) != 0 {
			count++
			database.DB.Debug().Create(&newUsers)
		}
	}

	return c.JSON(&fiber.Map{
		"message":                   "successfully update",
		"number of transfered data": count,
	})

}

//-------end of transfering data from another schema

//from another schema-------------------------------------------

func Return2(c *fiber.Ctx) error {

	var schema, table string
	schema = "mfs"
	table = "user"
	ReturnCtx(c, schema, table)
	return nil
}

func ReturnCtx(c *fiber.Ctx, schema, table string) error {

	var user []User
	database.DB.Debug().Raw(fmt.Sprintf("Select * from %s.%s", schema, table)).Scan(&user)

	fmt.Println(user)
	return c.JSON(&fiber.Map{
		"message": "userss",
		"data":    user,
	})
}

//end of from another schema ---------------------------------------

func Unique(s []Users) []Users {
	inResult := make(map[string]bool)
	var result []Users
	for _, str := range s {
		if _, ok := inResult[str.Name]; !ok {
			inResult[str.Name] = true
			result = append(result, str)
		}
	}
	return result
}
