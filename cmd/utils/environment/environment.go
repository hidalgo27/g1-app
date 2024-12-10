package environment

import (
	"errors"
	"fmt"
	"github.com/hidalgo27/app-g1/cmd/utils/log"
	"os"
)

var logger = log.GetLogger()

var User string
var Password string
var Server string
var Database string
var Port string
var SSLMode string // Nueva variable para el modo SSL
var JwtSecret string

func ConfigVariables() error {
	User = os.Getenv("PG_USER")
	Password = os.Getenv("PG_PASSWORD")
	Server = os.Getenv("PG_HOST")
	Database = os.Getenv("PG_DATABASE")
	Port = os.Getenv("PG_PORT")
	SSLMode = os.Getenv("PG_SSLMODE") // Obtén el modo SSL
	JwtSecret = os.Getenv("JWT_SECRET")

	fmt.Println(JwtSecret)
	// Verificar si alguna variable de entorno no está definida
	if !validateVariables() {
		msg := "faltan definiciones de variables de entorno"
		logger.Error(msg)
		return errors.New(msg)
	}
	return nil
}

func validateVariables() bool {
	if User == "" || Password == "" || Server == "" || Database == "" || Port == "" || SSLMode == "" || JwtSecret == "" {
		return false
	}
	return true
}
