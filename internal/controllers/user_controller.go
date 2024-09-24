package controllers

import (
	"mini_blog/config"
	model "mini_blog/internal/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	user := model.User{ID: 1}

	dbResult := config.DB.First(&user)

	if dbResult.Error == gorm.ErrRecordNotFound {
		c.SendStatus(fiber.StatusNotFound)
		return c.JSON("Error not found")
	}

	return c.JSON(user)
}
