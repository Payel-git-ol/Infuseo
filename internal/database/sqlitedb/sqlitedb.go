package sqlitedb

import (
	"Infuseo/internal/registretion/User"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitSqlite() error {
	var err error
	// Подключаемся к SQLite (файл создастся автоматически)
	Db, err = gorm.Open(sqlite.Open("UserLite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка создания SQLITE")
	}

	// Автомиграция (создаёт таблицу по структуре)
	err = Db.AutoMigrate(&User.User{})
	if err != nil {
		log.Fatalf("Ошибка миграции структуры в SQLITE")
	}

	return nil
}
