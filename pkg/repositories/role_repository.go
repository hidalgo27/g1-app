package repositories

import "github.com/hidalgo27/app-g1/cmd/database"

// import "github.com/hidalgo27/app-g1/cmd/database"
func GetUserRoles(userID uint) ([]string, error) {
	var roles []string
	err := database.DB.Table("roles").
		Select("roles.name").
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}
