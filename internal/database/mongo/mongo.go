package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"Infuseo/internal/registretion/User" // Импортируем вашу структуру User

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

const (
	DATABASE_NAME    = "mongo"
	COLLECTION_USERS = "Users"
)

func init() {
	InitMongo()
}

func InitMongo() {
	mongoURI := "mongodb://localhost:27017"
	log.Printf("Попытка подключения к MongoDB по URI: %s", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("КРИТИЧЕСКАЯ ОШИБКА: Не удалось создать клиент MongoDB: %v", err)
	}
	log.Println("Клиент MongoDB создан.")

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("КРИТИЧЕСКАЯ ОШИБКА: Не удалось подключиться к MongoDB: %v", err)
	}
	log.Println("Успешно подключено к MongoDB.")

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		log.Fatalf("КРИТИЧЕСКАЯ ОШИБКА: Не удалось пингануть MongoDB: %v", err)
	}
	log.Println("Пинг MongoDB успешен.")

}

func GetMongoClient() *mongo.Client {
	if client == nil {
		log.Fatal("КРИТИЧЕСКАЯ ОШИБКА: Клиент MongoDB не инициализирован. Убедитесь, что InitMongo() был вызван.")
	}
	return client
}

func CloseMongoConnection() {
	if client != nil {
		disconnectCtx, disconnectCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer disconnectCancel()
		if err := client.Disconnect(disconnectCtx); err != nil {
			log.Printf("ОШИБКА ОТКЛЮЧЕНИЯ: %v", err)
		} else {
			log.Println("Успешно отключено от MongoDB!")
		}
	}
}

func InsertUser(user User.User) (*mongo.InsertOneResult, error) {
	mongoClient := GetMongoClient()
	collection := mongoClient.Database(DATABASE_NAME).Collection(COLLECTION_USERS)

	// Устанавливаем _id, если он не установлен, MongoDB сгенерирует его автоматически
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	// Устанавливаем CreatedAt, если не установлено
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("ошибка при вставке пользователя: %w", err)
	}
	log.Printf("Пользователь '%s' успешно вставлен в MongoDB с ID: %s", user.Username, result.InsertedID)
	return result, nil
}

// FindUserByEmailOrUsername ищет пользователя по email или username
func FindUserByEmailOrUsername(email, username string) (*User.User, error) {
	mongoClient := GetMongoClient()
	collection := mongoClient.Database(DATABASE_NAME).Collection(COLLECTION_USERS)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"username": username},
		},
	}

	var foundUser User.User
	err := collection.FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Пользователь не найден
		}
		return nil, fmt.Errorf("ошибка при поиске пользователя: %w", err)
	}
	return &foundUser, nil
}
