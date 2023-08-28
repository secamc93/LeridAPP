package controller

import (
	"LeridAPP/db"
	"LeridAPP/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controladores para CategoriaNegocio
func GetCategoriasNegocios(w http.ResponseWriter, r *http.Request) {
	var categoriasNegocios []models.CategoriaNegocio
	err := models.GetAllCategoriasNegocios(db.DBConn, &categoriasNegocios)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoriasNegocios)
}

func GetCategoriaNegocio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	categoriaNegocio, err := models.GetCategoriaNegocio(db.DBConn, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoriaNegocio)
}

func CreateCategoriaNegocio(w http.ResponseWriter, r *http.Request) {
	var categoriaNegocio models.CategoriaNegocio
	json.NewDecoder(r.Body).Decode(&categoriaNegocio)

	err := models.CreateCategoriaNegocio(db.DBConn, &categoriaNegocio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoriaNegocio)
}

func UpdateCategoriaNegocio(w http.ResponseWriter, r *http.Request) {
	var categoriaNegocio models.CategoriaNegocio
	json.NewDecoder(r.Body).Decode(&categoriaNegocio)

	err := models.UpdateCategoriaNegocio(db.DBConn, &categoriaNegocio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoriaNegocio)
}

func DeleteCategoriaNegocio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	err := models.DeleteCategoriaNegocio(db.DBConn, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("CategoriaNegocio eliminada")
}
