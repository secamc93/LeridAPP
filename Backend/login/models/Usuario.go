package models

import (
	"errors"
	"login/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre        string `gorm:"type:varchar(100)"`
	Usuario       string `gorm:"type:varchar(100);unique_index"`
	Email         string `gorm:"type:varchar(100);unique_index"`
	Contrasena    string `gorm:"type:varchar(100)"`
	Administrador bool   // Campo booleano para administrador
	FkNegocio     uint
	Negocio       Negocio `gorm:"foreignKey:FkNegocio"`
}

// GetUsuarioByUsuario busca un usuario por nombre de usuario en la base de datos
func GetUsuarioByUsuario(usuario *Usuario, nombreUsuario string) (err error) {
	if err = db.DBConn.Where("usuario = ?", nombreUsuario).First(usuario).Error; err != nil {
		return err
	}
	return nil
}
func CheckPassword(usuario *Usuario, contrasena string) error {
	// Comprueba si la contraseña proporcionada coincide con la contraseña almacenada.
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(contrasena))

	if err != nil {
		// Si hay un error, la contraseña no coincide
		return errors.New("contraseña incorrecta")
	}

	// No hubo error, así que las contraseñas coinciden
	return nil
}
