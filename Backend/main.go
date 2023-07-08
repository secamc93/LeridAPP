package main

import (
	"fmt"
	"LeridAPP/db"
	"LeridAPP/models"
	"LeridAPP/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	r := router.Router()

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println(tpl)
		}
		return nil
	})

	fmt.Println("Serving on port 8000")

	// Ejecutar AutoMigrate para crear las tablas en la base de datos
	err = db.DBConn.AutoMigrate(&models.Usuario{}, &models.Categoria{}, &models.Negocio{}, &models.NegocioCategoria{}, &models.Comentario{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
