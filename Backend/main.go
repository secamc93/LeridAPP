package main

import (
	"LeridAPP/db"
	"LeridAPP/models"
	"LeridAPP/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	err = db.DBConn.AutoMigrate(&models.Usuario{}, &models.Categoria{}, &models.Negocio{}, &models.Comentario{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	// Create a cors wrapper (middleware)
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
	})

	err = http.ListenAndServe(":8000", corsWrapper.Handler(r))
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
