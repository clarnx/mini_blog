package main

import (
	"fmt"
	"log"
	"mini_blog/config"
	"mini_blog/internal/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// Load env variables
	var (
		serverPort = os.Getenv("PORT")
		dbName     = os.Getenv("DB_DATABASE")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPort     = os.Getenv("DB_PORT")
		dbHost     = os.Getenv("DB_HOST")
	)

	config.ConnectDB(dbHost, dbUsername, dbPassword, dbName, dbPort)

	app := fiber.New()

	routes.RegisterApiRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", serverPort)))

	// gentool -db "postgres" -dsn "host=localhost user=postgres password=postgres dbname=edeybe_bnpl port=5432" -outPath "./database/models"
}
