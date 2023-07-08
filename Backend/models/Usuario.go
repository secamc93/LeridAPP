package models

import (
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	gorm.Model
	Nombre        string
	Email         string `gorm:"type:varchar(100);unique_index"`
	Contrasena    string
	FechaRegistro time.Time
	Negocios      []Negocio
	Comentarios   []Comentario
}

func GetAllUsuarios(db *gorm.DB, usuarios *[]Usuario) (err error) {
	if err = db.Preload("Negocios").Preload("Comentarios").Find(usuarios).Error; err != nil {
		return err
	}
	return nil
}

func GetUsuario(db *gorm.DB, usuario *Usuario, id string) (err error) {
	if err = db.Where("id = ?", id).First(usuario).Error; err != nil {
		return err
	}
	return nil
}

func CrearUsuario(db *gorm.DB, usuario *Usuario) (err error) {
	if err = db.Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func ActualizarUsuario(db *gorm.DB, usuario *Usuario) (err error) {
	db.Save(usuario)
	return nil
}

func EliminarUsuario(db *gorm.DB, id string) (err error) {
	db.Where("id = ?", id).Delete(Usuario{})
	return nil
}
