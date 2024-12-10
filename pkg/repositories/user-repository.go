package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}
