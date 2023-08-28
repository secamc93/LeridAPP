package models

import (
	"gorm.io/gorm"
	"time"
)

type Comentario struct {
	gorm.Model
	Comentario string
	Valoracion int
	Fecha      time.Time
	UsuarioID  uint
	NegocioID  uint
}

// Obtén todos los comentarios
func GetAllComentarios(db *gorm.DB, comentarios *[]Comentario) (err error) {
	if err = db.Find(comentarios).Error; err != nil {
		return err
	}
	return nil
}

// Obtén un comentario por ID
func GetComentario(db *gorm.DB, comentario *Comentario, id string) (err error) {
	if err = db.Where("id = ?", id).First(comentario).Error; err != nil {
		return err
	}
	return nil
}

// Crea un nuevo comentario
func CrearComentario(db *gorm.DB, comentario *Comentario) (err error) {
	if err = db.Create(comentario).Error; err != nil {
		return err
	}
	return nil
}

// Actualiza un comentario existente
func ActualizarComentario(db *gorm.DB, comentario *Comentario) (err error) {
	db.Save(comentario)
	return nil
}

// Elimina un comentario
func EliminarComentario(db *gorm.DB, id string) (err error) {
	db.Where("id = ?", id).Delete(&Comentario{})
	return nil
}
