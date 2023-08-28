package main

import (
	"fmt"
	"log"
	"login/db"
	"login/models"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Usuario{}) // Luego, la tabla "usuario"
	if err != nil {
		log.Fatalf("Failed to migrate table Usuario: %v", err)
	}

	fmt.Println("Migration successful")
}
