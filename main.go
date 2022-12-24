package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/raihanbachdim/ProjectSIU/db"
	"github.com/raihanbachdim/ProjectSIU/routes"
)

func main() {
	db.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak bisa mengunduh file env nya bang")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
