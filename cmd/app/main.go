package main

import (
	mongodb "Infuseo/internal/database/mongo"
	"Infuseo/internal/registretion"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	mongodb.InitMongo()

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./web/static")

	app.Get("/reg", registration.GetHandlerRegister)
	app.Post("/reg", registration.PostHandlerRegister)

	app.Listen(":8080")
}
