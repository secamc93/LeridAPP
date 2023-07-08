package router

import (
	"github.com/gorilla/mux"
	"LeridAPP/controller" 
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Rutas para Usuarios
	r.HandleFunc("/api/usuarios", controller.GetUsuarios).Methods("GET")
	r.HandleFunc("/api/usuarios/{id}", controller.GetUsuario).Methods("GET")
	r.HandleFunc("/api/usuarios", controller.CrearUsuario).Methods("POST")
	r.HandleFunc("/api/usuarios/{id}", controller.ActualizarUsuario).Methods("PUT")
	r.HandleFunc("/api/usuarios/{id}", controller.EliminarUsuario).Methods("DELETE")

	// Rutas para Categorías
	r.HandleFunc("/api/categorias", controller.GetCategorias).Methods("GET")
	r.HandleFunc("/api/categorias/{id}", controller.GetCategoria).Methods("GET")
	r.HandleFunc("/api/categorias", controller.CrearCategoria).Methods("POST")
	r.HandleFunc("/api/categorias/{id}", controller.ActualizarCategoria).Methods("PUT")
	r.HandleFunc("/api/categorias/{id}", controller.EliminarCategoria).Methods("DELETE")

	// Rutas para Negocios
	r.HandleFunc("/api/negocios", controller.GetNegocios).Methods("GET")
	r.HandleFunc("/api/negocios/{id}", controller.GetNegocio).Methods("GET")
	r.HandleFunc("/api/negocios", controller.CrearNegocio).Methods("POST")
	r.HandleFunc("/api/negocios/{id}", controller.ActualizarNegocio).Methods("PUT")
	r.HandleFunc("/api/negocios/{id}", controller.EliminarNegocio).Methods("DELETE")

	// Rutas para Comentarios
	r.HandleFunc("/api/comentarios", controller.GetComentarios).Methods("GET")
	r.HandleFunc("/api/comentarios/{id}", controller.GetComentario).Methods("GET")
	r.HandleFunc("/api/comentarios", controller.CrearComentario).Methods("POST")
	r.HandleFunc("/api/comentarios/{id}", controller.ActualizarComentario).Methods("PUT")
	r.HandleFunc("/api/comentarios/{id}", controller.EliminarComentario).Methods("DELETE")

	return r
}