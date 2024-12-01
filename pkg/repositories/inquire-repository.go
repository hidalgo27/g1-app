package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func InquireCount(total int64) (int64, error) {
	err := database.DB.Model(&models.Inquire{}).Count(&total).Error
	return total, err
}

func Inquire(offset int, size int, inquires []models.Inquire) ([]models.Inquire, error) {
	err := database.DB.Model(&models.Inquire{}).Offset(offset).Limit(size).Find(&inquires).Error
	return inquires, err
}
