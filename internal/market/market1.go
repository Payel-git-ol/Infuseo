package market

import (
	"Infuseo/internal/database/postgresdb"
	"Infuseo/internal/database/postgresdb/productmodel"
	"github.com/gofiber/fiber/v2"
)

func GetHandlerMarket1(c *fiber.Ctx) error {
	var products []productmodel.Product

	if err := postgresdb.Db.Order("id desc").Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Title": "Ошибка",
			"Error": "Не удалось загрузить товары",
		})
	}

	return c.Render("market1", fiber.Map{"Products": products})
}

func PostHandlerMarket1(c *fiber.Ctx) error {
	return nil
}
