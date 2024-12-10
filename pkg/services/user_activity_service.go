package services

import "github.com/hidalgo27/app-g1/pkg/repositories"

func LogUserActivity(userID uint, action string) error {
	return repositories.LogUserActivity(userID, action)
}
