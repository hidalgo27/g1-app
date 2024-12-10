package services

import (
	"fmt"
	"github.com/hidalgo27/app-g1/pkg/entities"
	"github.com/hidalgo27/app-g1/pkg/mappers"
	"github.com/hidalgo27/app-g1/pkg/models"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

func GetInquire(page, size int) ([]entities.InquireDTO, int64, error) {
	var inquires []models.Inquire

	//var total int64
	offset := (page - 1) * size
	total, err := repositories.InquireCount()

	fmt.Println(total)

	if err != nil {
		return nil, 0, err
	}

	inquires, err = repositories.Inquire(offset, size)
	fmt.Println(inquires)
	if err != nil {
		return nil, 0, err
	}

	inquireResponses := mappers.MapInquireDTO(inquires)

	return inquireResponses, total, nil
}
