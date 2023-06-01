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

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	// cek data ada / tidak
	userId := ctx.Params("id")
	var user entity.User
	result := database.DB.Debug().First(&user, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}
	// parsing raw json
	p := new(request.UserCreateRequest)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	errors := ValidateStruct(*p)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"data":   errors,
		})

	}
	user.Name = p.Name
	user.Email = p.Email
	user.Address = p.Address
	user.Phone = p.Phone
	updateData := database.DB.Debug().Save(&user)
	if updateData.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": "failed",
			"data":   result.Error.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user,
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
