package initializers

import (
	"go-rest-api/pkg/models"
	"log"
)

func SyncDB() {
	DB.AutoMigrate(&models.Post{})

	log.Println("Database synced")
}
