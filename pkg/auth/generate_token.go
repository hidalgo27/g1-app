package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hidalgo27/app-g1/config"

	//"github.com/hidalgo27/app-g1/pkg/services"

	"time"
)

type Claims struct {
	UserID      uint     `json:"user_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateToken genera un token JWT con UserID y expiración
func GenerateToken(userID uint, roles []string, permissions []string) (string, error) {

	// Configuración de claims
	claims := Claims{
		UserID:      userID,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira en 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Crear el token con los claims y firmarlo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}
