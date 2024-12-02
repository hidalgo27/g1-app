package repositories

import (
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func GetPackagesWithPrices(packages *[]models.TPaquetes) error {
	err := database.DB.Preload("TPrecioPaquetes").Find(packages).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPackagesWithItineraries(packages *[]models.TPaquetes) error {
	err := database.DB.
		Preload("TPaqueteItinerarios.TItinerario"). // Carga itinerarios relacionados
		Find(packages).Error
	return err
}
