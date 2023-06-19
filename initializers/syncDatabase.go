package initializers

import (
	"github.com/drsimplegraffiti/gojwt/models"
)


func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}