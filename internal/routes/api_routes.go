package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")

	RegisterAuthRoutes(apiRouter)
	RegisterUserRoutes(apiRouter)
	RegisterPostRoutes(apiRouter)

}
