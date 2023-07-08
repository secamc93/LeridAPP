package models

import (
	"gorm.io/gorm"
)

type NegocioCategoria struct {
	gorm.Model
	NegocioID   uint
	CategoriaID uint
}

func GetAllNegocioCategorias(db *gorm.DB, negocioCategorias *[]NegocioCategoria) (err error) {
	if err = db.Find(negocioCategorias).Error; err != nil {
		return err
	}
	return nil
}

func GetNegocioCategoria(db *gorm.DB, negocioCategoria *NegocioCategoria, id string) (err error) {
	if err = db.Where("id = ?", id).First(negocioCategoria).Error; err != nil {
		return err
	}
	return nil
}

func CrearNegocioCategoria(db *gorm.DB, negocioCategoria *NegocioCategoria) (err error) {
	if err = db.Create(negocioCategoria).Error; err != nil {
		return err
	}
	return nil
}

func ActualizarNegocioCategoria(db *gorm.DB, negocioCategoria *NegocioCategoria) (err error) {
	db.Save(negocioCategoria)
	return nil
}

func EliminarNegocioCategoria(db *gorm.DB, id string) (err error) {
	db.Where("id = ?", id).Delete(NegocioCategoria{})
	return nil
}
