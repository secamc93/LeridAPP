package controller

import (
	"encoding/json"
	"net/http"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"github.com/gorilla/mux"
	"LeridAPP/db"
)


// Controladores para Negocios
func GetNegocios(w http.ResponseWriter, r *http.Request) {
	var negocios []models.Negocio
	err := models.GetAllNegocios(db.DBConn, &negocios) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocios)
}

func GetNegocio(w http.ResponseWriter, r *http.Request) {
	var negocio models.Negocio
	params := mux.Vars(r)
	err := models.GetNegocio(db.DBConn, &negocio, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocio)
}

func CrearNegocio(w http.ResponseWriter, r *http.Request) {
	var negocio models.Negocio
	json.NewDecoder(r.Body).Decode(&negocio)
	err := models.CrearNegocio(db.DBConn, &negocio) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocio)
}

func ActualizarNegocio(w http.ResponseWriter, r *http.Request) {
	var negocio models.Negocio
	json.NewDecoder(r.Body).Decode(&negocio)
	err := models.ActualizarNegocio(db.DBConn, &negocio) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(negocio)
}

func EliminarNegocio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.EliminarNegocio(db.DBConn, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Negocio eliminado")
}
