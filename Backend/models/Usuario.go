package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre      string
	Email       string `gorm:"type:varchar(100);unique_index"`
	Contrasena  string
	Negocios    []Negocio    // Asegúrate de que este tipo coincida con tu modelo de Negocio
	Comentarios []Comentario // Asegúrate de que este tipo coincida con tu modelo de Comentario
}

func GetAllUsuarios(db *gorm.DB, usuarios *[]Usuario) (err error) {
	if err = db.Find(usuarios).Error; err != nil {
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuario.Contrasena = string(hashedPassword)
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
func IniciarSesion(db *gorm.DB, email, contrasena string) (usuario Usuario, err error) {
	if err = db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return usuario, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(contrasena))
	if err != nil {
		return usuario, err
	}

	return usuario, nil
}
