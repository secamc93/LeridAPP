package controller

import (
	"encoding/json"
	"net/http"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"github.com/gorilla/mux"
	"LeridAPP/db"
)


// Controladores para Usuarios
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	err := models.GetAllUsuarios(db.DBConn, &usuarios) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuarios)
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	params := mux.Vars(r)
	err := models.GetUsuario(db.DBConn, &usuario, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	err := models.CrearUsuario(db.DBConn, &usuario) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	err := models.ActualizarUsuario(db.DBConn, &usuario) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.EliminarUsuario(db.DBConn, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Usuario eliminado")
}

