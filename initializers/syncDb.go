package initializers

import (
	"go-rest-api/models"
	"log"
)

func SyncDB() {
	DB.AutoMigrate(&models.Post{})

	log.Println("Database synced")
}
