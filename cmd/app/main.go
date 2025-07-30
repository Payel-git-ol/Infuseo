package main

import (
	"Infuseo/internal/database/sqlitedb"
	"Infuseo/internal/registretion"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	// Инициализация БД SQLITE
	if err := sqlitedb.InitSqlite(); err != nil {
		log.Fatalf("Ошибка инициализации БД SQLITE: %v", err)
	}

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./web/static")

	app.Get("/reg", registration.GetHandlerRegister)
	app.Post("/reg", registration.PostHandlerRegister)

	log.Fatal(app.Listen(":8080"))
}
