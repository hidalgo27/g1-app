package interfaces

type ProductRepository interface {
	//UserHasAccessToProduct(userID uint, productName string) (bool, error)
	UserHasAccessToProductAndRole(userID uint, productName string) (bool, error)
}
