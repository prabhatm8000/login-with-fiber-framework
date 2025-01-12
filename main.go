package main

import (
	"log"

	"example.com/login/configs/mongodb"
	"example.com/login/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found.")
	}

	mongodb.ConnectMongoDB()
	app := fiber.New()
	routes.SetupAPIRoutes(app)

	app.Static("/", "./UI")

	app.Listen(":3000")
}
