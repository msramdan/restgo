package handler

import (
	"rest-go/database"
	"rest-go/model/entity"
	"rest-go/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}
	return ctx.JSON(users)
}
func UserHandlerCreate(ctx *fiber.Ctx) error {
	// get dt raw body
	p := new(request.UserCreateRequest)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	// validator
	errors := ValidateStruct(*p)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"data":    errors,
		})

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
			"message": "failed",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

var validate = validator.New()

func ValidateStruct(p request.UserCreateRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
