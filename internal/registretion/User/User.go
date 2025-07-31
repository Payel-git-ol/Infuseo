package User

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // <-- ЭТОТ ИМПОРТ КРИТИЧЕН
)

// User представляет собой пользователя в MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // <-- ЭТОТ ТИП КРИТИЧЕН
	Username  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,omitempty"` // Хешированный пароль
	Email     string             `bson:"email,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
}
