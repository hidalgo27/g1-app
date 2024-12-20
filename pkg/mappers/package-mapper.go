package mappers

import (
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func MapPackageDTO(packages []models.TPaquetes) []entities.PackageDTO {
	packageResponses := make([]entities.PackageDTO, len(packages))

	for i, p := range packages {
		var prices []entities.PackagePriceDTO

		for _, price := range p.TPrecioPaquetes {
			prices = append(prices, entities.PackagePriceDTO{
				ID:        price.ID,
				Estrellas: price.Estrellas,
				PrecioS:   price.PrecioS,
				PrecioD:   price.PrecioD,
				PrecioT:   price.PrecioT,
				CodigoS:   price.CodigoS,
				CodigoD:   price.CodigoD,
				CodigoT:   price.CodigoT,
			})
		}

		packageResponses[i] = entities.PackageDTO{
			ID:          p.ID,
			Codigo:      p.Codigo,
			Titulo:      p.Titulo,
			PrecioTours: p.PrecioTours,
			Url:         p.URL,
			Prices:      prices,
		}
	}

	return packageResponses
}

func MapPackageWithItinerariesDTO(packages []models.TPaquetes) []entities.PackageWithItinerariesDTO {
	packageResponses := make([]entities.PackageWithItinerariesDTO, len(packages))

	for i, p := range packages {

		// Mapear itinerarios
		itineraries := MapItineraries(p.TPaqueteItinerarios)

		// Mapear el paquete principal
		packageResponses[i] = entities.PackageWithItinerariesDTO{
			ID:          p.ID,
			Codigo:      p.Codigo,
			Titulo:      p.Titulo,
			PrecioTours: p.PrecioTours,
			Url:         p.URL,
			Itineraries: itineraries,
		}
	}

	return packageResponses
}
