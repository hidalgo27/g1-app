package database

import (
	"fmt"
	"github.com/hidalgo27/app-g1/cmd/utils/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	LoadEnv()

	// Cambiamos el DSN para PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		environment.Server,
		environment.User,
		environment.Password,
		environment.Database,
		environment.Port,
		environment.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func InitDB() {
	db := ConnectToDB()
	DB = db
	CleanExpiredTokens()
}
