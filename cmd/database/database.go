package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hidalgo27/app-g1/cmd/utils/environment"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func loadEnv() {
	env := os.Getenv("GO_ENV")
	var envFile string

	switch env {
	case "dev":
		envFile = ".env.dev"
	case "test":
		envFile = ".env.test"
	default:
		envFile = ".env"
	}

	fmt.Println(envFile)

	errorEnv := godotenv.Load(envFile)
	if errorEnv != nil {
		log.Fatal(errorEnv)
	}

	errVar := environment.ConfigVariables()
	if errVar != nil {
		log.Fatal(errVar)
	}
}

func ConnectToDB() *gorm.DB {
	loadEnv()
	dsn := "" + environment.User + ":" + environment.Password + "@tcp(" + environment.Server + ":" + environment.Port + ")/" + environment.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitDB() {
	db := ConnectToDB()
	DB = db
}
