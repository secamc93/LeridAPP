// migrate.go
package main

import (
	"LeridAPP/db"
	"LeridAPP/models"
	"fmt"
	"log"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Usuario{}, &models.Categoria{}, &models.Negocio{}, &models.Comentario{}, &models.CategoriaNegocio{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	fmt.Println("Migration successful")
}
