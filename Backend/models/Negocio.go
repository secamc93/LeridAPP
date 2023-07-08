package models

import (
	"gorm.io/gorm"
)

type Negocio struct {
	gorm.Model
	Nombre          string
	Direccion       string
	Telefono        string
	Horario         string
	Descripcion     string
	UrlImagen       string
	UsuarioID       uint
	NegociosCategorias []NegocioCategoria
	Comentarios     []Comentario
}

func GetAllNegocios(db *gorm.DB, negocios *[]Negocio) (err error) {
	if err = db.Preload("NegociosCategorias").Preload("Comentarios").Find(negocios).Error; err != nil {
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
	db.Where("id = ?", id).Delete(Negocio{})
	return nil
}
