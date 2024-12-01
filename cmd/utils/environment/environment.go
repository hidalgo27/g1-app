package environment

import (
	"errors"
	"github.com/hidalgo27/app-g1/cmd/utils/log"
	"os"
)

var logger = log.GetLogger()

var User string
var Password string
var Server string
var Database string
var Port string

func ConfigVariables() error {
	User = os.Getenv("MYSQL_USER")
	Password = os.Getenv("MYSQL_PASSWORD")
	Server = os.Getenv("MYSQL_SERVER")
	Database = os.Getenv("MYSQL_DATABASE")
	Port = os.Getenv("MYSQL_PORT")

	// Verificar si alguna variable de entorno no está definida
	if !validateVariables() {
		msg := "faltan definiciones de variables de entorno"
		logger.Error(msg)
		return errors.New(msg)
	}
	return nil
}

func validateVariables() bool {
	// Descomentar y validar las variables necesarias
	if User == "" || Password == "" || Server == "" || Database == "" || Port == "" {
		return false
	}
	return true
}
