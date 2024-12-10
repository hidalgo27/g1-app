package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func LogUserActivity(userID uint, action string) error {
	activity := models.UserActivity{
		UserID: userID,
		Action: action,
	}

	result := database.DB.Create(&activity)
	return result.Error
}
