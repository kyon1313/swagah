package route

import (
	"fmt"
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

type Errror struct {
	ErrorMessage string `json:"errormessage" `
}

// Add User godoc
// @Summary     Add users
// @Description  get string by ID
// @Tags         Adding
// @Accept       json
// @Produce      json
// @Param        user body model.User true "Add user"
// @Success      200  {object}  model.User
// @Failure      400  {object}  Errror
// @Router       /proj1/add/ [post]
func Add(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "internal server error",
			"model":   Errror{ErrorMessage: "error"},
		})
	}

	database.DB.Debug().Create(&user)

	return c.JSON(&fiber.Map{
		"data": user,
	})
}

// Show user godoc
// @Summary      show 1 user
// @Description  get string by ID
// @Tags         Getting
// @Accept       json
// @Produce      json
// @Param        id path int true "id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  Errror
// @Router       /proj1/get/{id} [get]
func Get(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"message": "Not found",
			"error":   Errror{ErrorMessage: "not fuckingfound"},
		})
	}
	return c.JSON(&fiber.Map{
		"message": "success",
		"date":    user,
	})

}

// Show all users godoc
// @Summary      Show all users
// @Description  get string by ID
// @Tags         Getting
// @Produce      json
// @Success      200  {object}  model.User
// @Failure      400  {object}  Errror
// @Router       /proj1/gets/ [get]
func GetAll(c *fiber.Ctx) error {
	var user []model.User
	// var users []User
	database.DB.Find(&user)
	if len(user) == 0 {
		return c.JSON(Errror{ErrorMessage: "erro"})
	}

	// for _, v := range user {
	// 	response := JustGetName(v)
	// 	users = append(users, response)
	// }

	return c.JSON(user)
}

func RecoverFromPanics(c *fiber.Ctx, u model.User) error {
	if r := recover(); r != nil {
		return c.JSON(&fiber.Map{
			"from panic": u,
		})
	}
	return nil
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Savings  string `json:"savings"`
}

func JustGetName(user model.User) User {
	return User{
		Name: user.Name,
	}
}

// Show user godoc
// @Summary      Delete user
// @Description  get string by ID
// @Tags         Deleting
// @Accept       json
// @Produce      json
// @Param        id path int true "id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  Errror
// @Router       /proj1/delete/{id} [delete]
func Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"messagge": "Cannot find",
		})
	}

	database.DB.Debug().Delete(&user)
	return c.JSON(&fiber.Map{
		"message": "successfully deleted",
		"success": true,
	})

}

// Add User godoc
// @Summary     Add users
// @Description  get string by ID
// @Tags         Updating
// @Accept       json
// @Produce      json
// @Param        user body model.User true "Update User"
// @Success      200  {object}  model.User
// @Failure      400  {object}  Errror
// @Router       /proj1/update/ [put]
func Update(c *fiber.Ctx) error {

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "internal server error",
			"model":   Errror{ErrorMessage: "error"},
		})
	}
	database.DB.Debug().Where("id=?", user.ID).Updates(&user)

	fmt.Println(user)
	return c.JSON(&fiber.Map{
		"message": "success",
		"date":    user,
	})

}

// type Error struct {
// 	Retcode      string `json:"retcode"`
// 	ErrorMessage string `json:"errorMEssage"`
// }

// func Errorsaa(e Error, retcode, em string) Error {
// 	return Error{
// 		Retcode:      retcode,
// 		ErrorMessage: em,
// 	}
// }
