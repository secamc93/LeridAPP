package controller

import (
	"encoding/json"
	"net/http"
	"LeridAPP/models" // Asegúrate de cambiar "LeridAPP" por el nombre de tu paquete
	"github.com/gorilla/mux"
	"LeridAPP/db"
)



// Controladores para Comentarios
func GetComentarios(w http.ResponseWriter, r *http.Request) {
	var comentarios []models.Comentario
	err := models.GetAllComentarios(db.DBConn, &comentarios) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(comentarios)
}

func GetComentario(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	params := mux.Vars(r)
	err := models.GetComentario(db.DBConn, &comentario, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(comentario)
}

func CrearComentario(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	json.NewDecoder(r.Body).Decode(&comentario)
	err := models.CrearComentario(db.DBConn, &comentario) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(comentario)
}

func ActualizarComentario(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	json.NewDecoder(r.Body).Decode(&comentario)
	err := models.ActualizarComentario(db.DBConn, &comentario) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(comentario)
}

func EliminarComentario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.EliminarComentario(db.DBConn, params["id"]) // Agrega db como argumento
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Comentario eliminado")
}
