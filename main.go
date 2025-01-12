package main

import (
	"log"
	"os"

	"example.com/login/configs/mongodb"
	"example.com/login/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(".env file not found.")
		}
	} else {
		log.Println("Running in PROD mode.")
	}

	mongodb.ConnectMongoDB()
	app := fiber.New()
	routes.SetupAPIRoutes(app)

	app.Static("/", "./UI")

	app.Listen(":3000")
}
