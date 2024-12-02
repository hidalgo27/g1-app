package mappers

import (
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func MapItineraries(paqueteItinerarios []models.TPaqueteItinerario) []entities.ItineraryDTO {
	itineraries := make([]entities.ItineraryDTO, len(paqueteItinerarios))

	//Iterar sobre los itinerarios asociados
	for i, itinerary := range paqueteItinerarios {
		itineraries[i] = entities.ItineraryDTO{
			ID:          itinerary.TItinerario.ID,
			Codigo:      itinerary.TItinerario.Codigo,
			Dia:         itinerary.TItinerario.Dia,
			Titulo:      itinerary.TItinerario.Titulo,
			Resumen:     itinerary.TItinerario.Resumen,
			Descripcion: itinerary.TItinerario.Descripcion,
		}
	}

	return itineraries
}
