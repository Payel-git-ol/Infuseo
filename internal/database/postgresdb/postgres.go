package postgresdb

import (
	"Infuseo/internal/database/postgresdb/productmodel"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func InitDbPostgres() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	dsn := os.Getenv("POSTGRES")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN not configured")
	}

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	log.Println("Successfully connected to PostgreSQL")

	// Миграция
	if err := Db.AutoMigrate(&productmodel.Product{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Загрузка данных
	/*if err := loadProducts(); err != nil {
		log.Fatal("Failed to load products:", err)
	}*/
}

/*func loadProducts() error {
	// 1. Чтение файла
	file, err := os.ReadFile("productmodel.json")
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// 2. Структура для парсинга
	type ProductList struct {
		Products []productmodel.Product `json:"products"`
	}

	var productList ProductList
	if err := json.Unmarshal(file, &productList); err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	if len(productList.Products) == 0 {
		return errors.New("файл не содержит товаров")
	}

	// 3. Включим логирование SQL
	dbWithLogger := Db.Session(&gorm.Session{
		Logger: Db.Logger.LogMode(logger.Info),
	})

	// 4. Оптимизированная вставка
	err = dbWithLogger.Transaction(func(tx *gorm.DB) error {
		// Пропускаем существующие записи
		clause := clause.OnConflict{DoNothing: true}

		batchSize := 100 // PostgreSQL легко обрабатывает 100+ записей за раз
		if err := tx.Clauses(clause).CreateInBatches(productList.Products, batchSize).Error; err != nil {
			return fmt.Errorf("ошибка пакетной вставки: %w", err)
		}
		return nil
	})

	if err != nil {
		// 5. Детальный анализ ошибки
		log.Printf("Первый продукт для анализа: %+v", productList.Products[0])

		// Проверка типа данных
		var testProduct productmodel.Product
		if err := Db.First(&testProduct).Error; err != nil {
			log.Printf("Проверка структуры таблицы: %v", err)
		}

		return fmt.Errorf("финальная ошибка вставки: %w", err)
	}

	// 6. Проверка результатов
	var count int64
	if err := Db.Model(&productmodel.Product{}).Count(&count).Error; err != nil {
		log.Printf("Ошибка проверки количества: %v", err)
	}
	log.Printf("Успешно загружено товаров: %d/%d", count, len(productList.Products))

	return nil
}*/
