package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Categoria struct {
	gorm.Model
	Nombre string
}

func GetAllCategorias(db *gorm.DB, categorias *[]Categoria) (err error) {
	if err = db.Find(categorias).Error; err != nil {
		return err
	}
	return nil
}

func GetCategoria(db *gorm.DB, categoria *Categoria, id string) (err error) {
	if err = db.Where("id = ?", id).First(categoria).Error; err != nil {
		return err
	}
	return nil
}

func CrearCategoria(db *gorm.DB, categoria *Categoria) (err error) {
	if err = db.Create(categoria).Error; err != nil {
		return err
	}
	return nil
}

func ActualizarCategoria(db *gorm.DB, categoria *Categoria) (err error) {
	db.Save(categoria)
	return nil
}

func EliminarCategoria(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).Delete(&Categoria{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
