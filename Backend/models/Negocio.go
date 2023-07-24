package models

import (
	"gorm.io/gorm"
)

type Negocio struct {
	gorm.Model
	Nombre      string
	Direccion   string
	Telefono    string
	Horario     string
	Descripcion string
	UrlImagen   string
	UsuarioID   uint
	Categorias  string
	Comentarios string
}

func GetAllNegocios(db *gorm.DB, negocios *[]Negocio) (err error) {
	if err = db.Find(negocios).Error; err != nil {
		return err
	}
	return nil
}

func GetNegocio(db *gorm.DB, negocio *Negocio, id string) (err error) {
	if err = db.Where("id = ?", id).First(negocio).Error; err != nil {
		return err
	}
	return nil
}

func CrearNegocio(db *gorm.DB, negocio *Negocio) (err error) {
	if err = db.Create(negocio).Error; err != nil {
		return err
	}
	return nil
}

func ActualizarNegocio(db *gorm.DB, negocio *Negocio) (err error) {
	db.Save(negocio)
	return nil
}

func EliminarNegocio(db *gorm.DB, id string) (err error) {
	negocio := &Negocio{} // Debes pasar un puntero a la estructura Negocio.
	db.Where("id = ?", id).Delete(negocio)
	return nil
}
