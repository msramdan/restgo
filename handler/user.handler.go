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

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
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
			"status": "failed",
			"data":   errors,
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
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users []entity.User
	result := database.DB.Debug().First(&users, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users []entity.User
	result := database.DB.Debug().First(&users, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}

	deleteData := database.DB.Debug().Unscoped().Delete(&users)
	if deleteData.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   deleteData.Error.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   "User was deleted",
	})
}
