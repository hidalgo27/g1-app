package services

import (
	"errors"
	"github.com/hidalgo27/app-g1/cmd/utils"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/mappers"
	"github.com/hidalgo27/app-g1/pkg/models"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

func CreateUser(userDTO dto.UserCreateDTO) (dto.UserDTO, error) {
	// Verificar si el correo ya existe
	existingUser, _ := repositories.FindUserByEmail(userDTO.Email)
	if existingUser != nil {
		return dto.UserDTO{}, errors.New("email is already in use")
	}

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		return dto.UserDTO{}, errors.New("failed to hash password")
	}

	// Crear el modelo de usuario
	user := models.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	// Guardar en la base de datos
	err = repositories.CreateUser(&user)
	if err != nil {
		return dto.UserDTO{}, errors.New("failed to create user")
	}

	// Mapear el usuario a un DTO para la respuesta
	userResponse := mappers.MapUserToDTO(&user)
	return userResponse, nil
}

func GetRolesAndPermissions(userID uint) ([]string, []string, error) {
	roles, err := repositories.GetUserRoles(userID)
	if err != nil {
		return nil, nil, err
	}

	permissions, err := repositories.GetUserPermissions(userID)
	if err != nil {
		return nil, nil, err
	}

	return roles, permissions, nil
}
