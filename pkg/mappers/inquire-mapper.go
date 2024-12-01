package mappers

import (
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func MapInquireDTO(inquires []models.Inquire) []entities.InquireDTO {
	inquireResponses := make([]entities.InquireDTO, len(inquires))
	for i, s := range inquires {
		inquireResponses[i] = entities.InquireDTO{
			ID:     s.ID,
			Hotel:  s.Hotel,
			Nombre: s.Nombre,
			Email:  s.Email,
		}
	}
	return inquireResponses
}
