package main

import (
	"github.com/hidalgo27/app-g1/cmd/utils"
	"log"

	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/pkg/models"
)

func main() {
	// Inicializa la conexión a la base de datos
	database.InitDB()

	// Define los datos del usuario
	name := "John Doe"
	email := "john.doe@example.com"
	password := "secret"

	// Hashea la contraseña
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// Crea el usuario
	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Guarda el usuario en la base de datos
	if err := database.DB.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	log.Println("User created successfully!")
}
