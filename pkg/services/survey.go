package services

import (
	"fmt"
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/mappers"
	"github.com/hidalgo27/app-g1/pkg/models"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

//type SurveyService struct {
//}
//
//func NewSurveyService() *SurveyService {
//	return &SurveyService{}
//}

func GetSurveys(page, size int) ([]entities.InquireDTO, int64, error) {
	var inquires []models.Inquire

	var total int64
	offset := (page - 1) * size
	total, err := repositories.InquireCount(total)

	fmt.Println(total)

	if err != nil {
		return nil, 0, err
	}

	inquires, err = repositories.Inquire(offset, size, inquires)
	fmt.Println(inquires)
	if err != nil {
		return nil, 0, err
	}

	inquireResponses := mappers.MapInquireDTO(inquires)

	return inquireResponses, total, nil
}

func GetInquire(page, size int) ([]entities.InquireDTO, int64, error) {
	var inquires []models.Inquire

	var total int64
	offset := (page - 1) * size
	total, err := repositories.InquireCount(total)

	fmt.Println(total)

	if err != nil {
		return nil, 0, err
	}

	inquires, err = repositories.Inquire(offset, size, inquires)
	fmt.Println(inquires)
	if err != nil {
		return nil, 0, err
	}

	inquireResponses := mappers.MapInquireDTO(inquires)

	return inquireResponses, total, nil
}

func translateStatus(status string) string {
	switch status {
	case "APPROVED":
		return "Aprobado"
	case "PENDING":
		return "Pendiente"
	case "PROCESS":
		return "En proceso"
	default:
		return "Desconocido" // Manejo de casos desconocidos o no mapeados
	}
}
