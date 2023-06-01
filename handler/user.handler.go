package handler

import (
	"rest-go/database"
	"rest-go/model/entity"
	"rest-go/model/request"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}
	return ctx.JSON(users)
}
func UserHandlerCreate(ctx *fiber.Ctx) error {

	p := new(request.UserCreateRequest)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	newUser := entity.User{
		Name:    p.Name,
		Email:   p.Email,
		Address: p.Address,
		Phone:   p.Phone,
	}
	result := database.DB.Debug().Create(&newUser)
	if result.Error != nil {
		return ctx.JSON(fiber.Map{
			"message": "Create User Failed",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Create User Success",
		"data":    newUser,
	})
}
