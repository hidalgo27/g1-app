package services

import (
	"errors"
	"github.com/hidalgo27/app-g1/cmd/utils"
	"github.com/hidalgo27/app-g1/pkg/auth"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/mappers"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

func Authenticate(email, password string) (string, string, dto.UserDTO, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("invalid email")
	}

	if !utils.CheckPassword(user.Password, password) {
		return "", "", dto.UserDTO{}, errors.New("invalid password")
	}

	roles, err := repositories.GetUserRoles(user.ID)
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("failed to fetch roles")
	}

	permissions, err := repositories.GetUserPermissions(user.ID)
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("failed to fetch permissions")
	}

	token, err := auth.GenerateToken(user.ID, roles, permissions)
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("failed to generate token")
	}

	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("failed to generate refresh token")
	}

	err = repositories.LogUserActivity(user.ID, "LOGIN")
	if err != nil {
		return "", "", dto.UserDTO{}, errors.New("failed to log user activity")
	}

	userDTO := mappers.MapUserToDTO(user)

	return token, refreshToken, userDTO, nil
}
