package main

import (
	mongodb "Infuseo/internal/database/mongo"
	"Infuseo/internal/database/postgresdb"
	"Infuseo/internal/handlerbuy"
	"Infuseo/internal/market"
	"Infuseo/internal/product"
	"Infuseo/internal/registretion"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	mongodb.InitMongo()
	postgresdb.InitDbPostgres()

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./web/static")

	app.Get("/reg", registration.GetHandlerRegister)
	app.Post("/reg", registration.PostHandlerRegister)

	app.Get("/market", market.GetHandlerMarket1)
	app.Post("/market", market.PostHandlerMarket1)

	app.Get("/product/:id", product.GetHandlerProduct)
	app.Post("/product/:id", product.PostHandlerProduct)

	app.Get("/buyproduct/:id", handlerbuy.GetBuy)

	app.Listen(":8080")
}
