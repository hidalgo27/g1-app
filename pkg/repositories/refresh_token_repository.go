package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
	"time"
)

func CreateRefreshToken(userID uint, token string, expiresAt time.Time) error {
	refreshToken := models.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	result := database.DB.Create(&refreshToken)
	return result.Error
}

func FindRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := database.DB.Where("token = ? AND expires_at > ?", token, time.Now()).First(&refreshToken).Error
	return &refreshToken, err

}
func DeleteRefreshToken(token string) error {
	return database.DB.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}

func DeleteAllRefreshTokensByUser(userID uint) error {
	result := database.DB.Where("user_id = ?", userID).Delete(&models.RefreshToken{})
	return result.Error
}
