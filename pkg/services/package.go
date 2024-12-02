package services

import (
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/mappers"
	"github.com/hidalgo27/app-g1/pkg/models"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

func GetAllPackagesWithPrices() ([]entities.PackageDTO, error) {
	var packages []models.TPaquetes

	// Obtener datos desde el repositories
	err := repositories.GetPackagesWithPrices(&packages)
	if err != nil {
		return nil, err
	}

	// Mapear datos
	packageDTOs := mappers.MapPackageDTO(packages)

	return packageDTOs, nil
}

func GetAllPackagesWithItineraries() ([]entities.PackageWithItinerariesDTO, error) {
	var packages []models.TPaquetes

	// Obtener datos desde el repositories
	err := repositories.GetPackagesWithItineraries(&packages)
	if err != nil {
		return nil, err
	}

	// Mapear datos
	packageDTOs := mappers.MapPackageWithItinerariesDTO(packages)

	return packageDTOs, nil
}
