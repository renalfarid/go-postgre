package main

import (
	"fmt"
	"log"

	"go-postgre/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnv()
	app := fiber.New()

	routes.SetupRoutes(app)
	// Register routes and start the server
	app.Listen(":3000")

	fmt.Println("Running Port :3000")
}
