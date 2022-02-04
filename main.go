package main

import (
	"fiber-cloudinary-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/file", controllers.FileUpload)
	app.Post("/remote", controllers.RemoteUpload)

	app.Listen(":6000")
}
