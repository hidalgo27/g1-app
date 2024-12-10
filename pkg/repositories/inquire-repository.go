package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func InquireCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Inquire{}).Count(&total).Error
	return total, err
}

func Inquire(offset int, size int) ([]models.Inquire, error) {
	var inquires []models.Inquire
	err := database.DB.Model(&models.Inquire{}).Offset(offset).Limit(size).Find(&inquires).Error
	return inquires, err
}
