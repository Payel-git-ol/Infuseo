package product

import (
	"Infuseo/internal/database/postgresdb"
	"Infuseo/internal/database/postgresdb/productmodel"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetHandlerProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product productmodel.Product

	if err := postgresdb.Db.First(&product, id).Error; err != nil {
		// Добавим логирование для отладки
		fmt.Printf("Ошибка при поиске товара ID %s: %v\n", id, err)
		return c.Status(404).SendString("Товар не найден")
	}

	return c.Render("product", fiber.Map{
		"Product": product,
	})
}

func PostHandlerProduct(c *fiber.Ctx) error {
	// Логика добавления в корзину
	return nil
}
