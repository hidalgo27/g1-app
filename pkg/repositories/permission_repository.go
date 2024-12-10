package repositories

import "github.com/hidalgo27/app-g1/cmd/database"

func GetUserPermissions(userID uint) ([]string, error) {
	var permissions []string
	err := database.DB.Table("permissions").
		Select("permissions.name").
		Joins("JOIN role_has_permissions ON permissions.id = role_has_permissions.permission_id").
		Joins("JOIN user_roles ON user_roles.role_id = role_has_permissions.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&permissions).Error
	return permissions, err
}
