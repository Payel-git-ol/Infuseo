package registration

import (
	"Infuseo/internal/database/mongo" // Изменяем импорт с sqlitedb на mongo
	"Infuseo/internal/registretion/User"
	"fmt"
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

	// Проверяем, существует ли пользователь в MongoDB
	existingUser, err := mongo.FindUserByEmailOrUsername(email, username)
	if err != nil {
		// Это ошибка не типа ErrNoDocuments, а какая-то другая ошибка БД
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Ошибка при проверке существования пользователя: %v", err),
		})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Пользователь с таким email или именем уже существует",
		})
	}

	// Создаем нового пользователя
	newUser := User.User{
		// ID будет сгенерирован автоматически MongoDB, т.к. omitempty
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: time.Now(),
	}

	// Сохраняем пользователя в MongoDB
	_, err = mongo.InsertUser(newUser) // Вызываем функцию из пакета mongo
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Ошибка при создании пользователя в базе данных: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		//"message": "Пользователь успешно зарегистрирован",
		// Здесь можно вернуть более подробную информацию, если нужно
		// "user": fiber.Map{
		//    "id":       result.InsertedID, // ID теперь из MongoDB
		//    "username": newUser.Username,
		//    "email":    newUser.Email,
		// },
	})
}
