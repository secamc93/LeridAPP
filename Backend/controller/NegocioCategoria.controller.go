package controller

import (
	"encoding/json"
	"net/http"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"github.com/gorilla/mux"
	"LeridAPP/db"
)


// Controladores para NegocioCategoria
func GetNegocioCategorias(w http.ResponseWriter, r *http.Request) {
	var negocioCategorias []models.NegocioCategoria
	err := models.GetAllNegocioCategorias(db.DBConn, &negocioCategorias) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocioCategorias)
}

func GetNegocioCategoria(w http.ResponseWriter, r *http.Request) {
	var negocioCategoria models.NegocioCategoria
	params := mux.Vars(r)
	err := models.GetNegocioCategoria(db.DBConn, &negocioCategoria, params["id"]) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocioCategoria)
}

func CrearNegocioCategoria(w http.ResponseWriter, r *http.Request) {
	var negocioCategoria models.NegocioCategoria
	json.NewDecoder(r.Body).Decode(&negocioCategoria)
	err := models.CrearNegocioCategoria(db.DBConn, &negocioCategoria) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocioCategoria)
}

func ActualizarNegocioCategoria(w http.ResponseWriter, r *http.Request) {
	var negocioCategoria models.NegocioCategoria
	json.NewDecoder(r.Body).Decode(&negocioCategoria)
	err := models.ActualizarNegocioCategoria(db.DBConn, &negocioCategoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocioCategoria)
}

func EliminarNegocioCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.EliminarNegocioCategoria(db.DBConn, params["id"]) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("NegocioCategoria eliminada")
}
