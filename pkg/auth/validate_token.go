package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hidalgo27/app-g1/config"
	"time"
)

// Claims estructura para almacenar los datos dentro del token JWT

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validar que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}

	// Extraer las claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Verificar si el token ha expirado
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
