package repositories

import "github.com/hidalgo27/app-g1/cmd/database"

// ProductRepositoryImpl implementa la interfaz ProductRepository
type ProductRepositoryImpl struct{}

// verifica si el usuario tiene acceso al producto
func (r *ProductRepositoryImpl) UserHasAccessToProduct(userID uint, productName string) (bool, error) {
	var count int64
	err := database.DB.Table("products").
		Joins("JOIN user_products ON products.id = user_products.product_id").
		Where("user_products.user_id = ? AND products.name = ?", userID, productName).
		Count(&count).Error

	return count > 0, err
}

func (r *ProductRepositoryImpl) UserHasAccessToProductAndRole(userID uint, productName string) (bool, error) {
	var count int64
	err := database.DB.Table("products").
		Joins("JOIN user_products ON products.id = user_products.product_id").
		Joins("JOIN user_roles ON user_products.user_id = user_roles.user_id").
		Joins("JOIN role_products ON role_products.product_id = products.id AND role_products.role_id = user_roles.role_id").
		Where("user_products.user_id = ? AND products.name = ?", userID, productName).
		Count(&count).Error

	return count > 0, err
}
