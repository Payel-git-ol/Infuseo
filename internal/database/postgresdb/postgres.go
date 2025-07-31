package postgresdb

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func init() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//InitDbPostgres()
	if err := InitDbPostgres; err != nil {
		log.Fatal("Ошибка инициализации БД POSTGRES")
	}
}

func InitDbPostgres() {
	dsn := os.Getenv("POSTGRES")

	if dsn == "" {
		log.Fatal("Ошибка получения данных с .env для POSTGRES")
	}
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД Postgres:", err)
	}
	fmt.Println("Успешное подключение к БД Postgres")
	// Здесь нужно указать ваши модели для миграции
	Db.AutoMigrate()
}
