package controller

import (
	"encoding/json"
	"net/http"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"github.com/gorilla/mux"
	"LeridAPP/db"
)



// Controladores para Categorias
func GetCategorias(w http.ResponseWriter, r *http.Request) {
	var categorias []models.Categoria
	err := models.GetAllCategorias(db.DBConn, &categorias) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categorias)
}

func GetCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	params := mux.Vars(r)
	err := models.GetCategoria(db.DBConn, &categoria, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	json.NewDecoder(r.Body).Decode(&categoria)
	err := models.CrearCategoria(db.DBConn, &categoria) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	json.NewDecoder(r.Body).Decode(&categoria)
	err := models.ActualizarCategoria(db.DBConn, &categoria) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.EliminarCategoria(db.DBConn, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Categoria eliminada")
}
