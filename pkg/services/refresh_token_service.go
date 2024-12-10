package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hidalgo27/app-g1/pkg/auth"
	"github.com/hidalgo27/app-g1/pkg/repositories"
	"time"
)

func GenerateRefreshToken(userID uint) (string, error) {
	token := uuid.New().String()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	err := repositories.CreateRefreshToken(userID, token, expiresAt)
	if err != nil {
		return "", errors.New("error creating refresh token")
	}
	return token, nil
}

func RefreshAccessToken(refreshToken string) (string, string, error) {
	storedToken, err := repositories.FindRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}

	roles, permission, err := GetRolesAndPermissions(storedToken.UserID)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}

	newAccessToken, err := auth.GenerateToken(storedToken.UserID, roles, permission)
	if err != nil {
		return "", "", errors.New("error generating new access token")
	}

	newRefreshToken, err := GenerateRefreshToken(storedToken.UserID)
	if err != nil {
		return "", "", errors.New("failed to generate new access token")
	}

	err = repositories.LogUserActivity(storedToken.UserID, "REFRESH_TOKEN")
	if err != nil {
		return "", "", errors.New("failed to log user activity")
	}

	err = repositories.DeleteRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("failed to delete refresh token")
	}
	return newAccessToken, newRefreshToken, nil
}
