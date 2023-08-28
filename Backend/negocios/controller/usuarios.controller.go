package controller

import (
	"LeridAPP/db"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	usuario.Contrasena = "" // Puedes borrar la contraseña cifrada antes de enviarla en la respuesta
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
func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email      string `json:"email"`
		Contrasena string `json:"contrasena"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuario, err := models.IniciarSesion(db.DBConn, input.Email, input.Contrasena)
	if err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"message": "Inicio de sesión exitoso",
		"user_id": usuario.ID,
	}
	json.NewEncoder(w).Encode(response)
}
