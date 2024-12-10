package mappers

import (
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func MapUserToDTO(user *models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
