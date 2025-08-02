package handlerbuy

import (
	"Infuseo/internal/database/postgresdb"
	"Infuseo/internal/database/postgresdb/productmodel"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetBuy(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).SendString("ID товара не указан")
	}

	var product productmodel.Product
	if err := postgresdb.Db.First(&product, id).Error; err != nil {
		fmt.Printf("Ошибка при поиске товара ID %s: %v\n", id, err)
		return c.Status(404).SendString("Товар не найден")
	}

	return c.Render("modelbuy", fiber.Map{
		"Product": product,
	})
}
