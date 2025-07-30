package registration

import (
	"Infuseo/internal/database/sqlitedb"
	"Infuseo/internal/registretion/User"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetHandlerRegister(c *fiber.Ctx) error {
	return c.Render("registration", fiber.Map{
		"Title": "Регистрация",
	})
}

func PostHandlerRegister(c *fiber.Ctx) error {
	// Получаем данные из формы
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	// Валидация
	if username == "" || password == "" || email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Все поля обязательны для заполнения",
		})
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при обработке пароля",
		})
	}

	// Проверяем, существует ли пользователь
	var existingUser User.User
	if err := sqlitedb.Db.Where("email = ? OR username = ?", email, username).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Пользователь с таким email или именем уже существует",
		})
	}

	// Создаем нового пользователя
	newUser := User.User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := sqlitedb.Db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при создании пользователя",
		})
	}

	return c.JSON(fiber.Map{
		/*"message": "Пользователь успешно зарегистрирован",
		"user": fiber.Map{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
		},*/
	})
}
