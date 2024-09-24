package controllers

import (
	"fmt"
	"mini_blog/config"
	model "mini_blog/internal/database/models"
	"mini_blog/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SignUpUser(c *fiber.Ctx) error {

	userData := new(types.UserSignUpData)

	if err := c.BodyParser(userData); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(500)
	}

	fmt.Println(userData)

	validate := validator.New()

	// errs := validate.Struct(userData)

	type ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	errs := validate.Struct(userData)
	validationErrors := []ErrorResponse{}

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}

		return c.JSON(validationErrors)
	}

	// if errs != nil {
	// 	fmt.Println(errs)
	// 	c.SendStatus(fiber.StatusBadRequest)
	// 	return c.JSON(errs)
	// }

	user := model.User{}

	dbResult := config.DB.First(&user)

	if dbResult.Error == gorm.ErrRecordNotFound {
		c.SendStatus(fiber.StatusNotFound)
		return c.JSON("Error not found")
	}

	return c.JSON(user)
}
