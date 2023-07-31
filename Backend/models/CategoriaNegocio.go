package models

import (
	"errors"

	"gorm.io/gorm"
)

type CategoriaNegocio struct {
	gorm.Model
	NegocioID   uint
	CategoríaID uint
	Negocio     Negocio   `gorm:"foreignkey:NegocioID"`
	Categoría   Categoria `gorm:"foreignkey:CategoríaID"`
}

// CreateCategoriaNegocio agrega una nueva relación entre una categoría y un negocio
func CreateCategoriaNegocio(db *gorm.DB, categoriaNegocio *CategoriaNegocio) error {
	return db.Create(categoriaNegocio).Error
}

// GetAllCategoriasNegocios obtiene todas las relaciones entre categorías y negocios
func GetAllCategoriasNegocios(db *gorm.DB, categoriaNegocios *[]CategoriaNegocio) error {
	return db.Preload("Negocio").Preload("Categoría").Find(categoriaNegocios).Error
}

// GetCategoriaNegocio obtiene la relación entre una categoría y un negocio por ID
func GetCategoriaNegocio(db *gorm.DB, id uint) (*CategoriaNegocio, error) {
	var categoriaNegocio CategoriaNegocio
	if err := db.Preload("Negocio").Preload("Categoría").First(&categoriaNegocio, id).Error; err != nil {
		return nil, err
	}
	return &categoriaNegocio, nil
}

// UpdateCategoriaNegocio actualiza una relación existente entre una categoría y un negocio
func UpdateCategoriaNegocio(db *gorm.DB, categoriaNegocio *CategoriaNegocio) error {
	if categoriaNegocio.ID == 0 {
		return errors.New("invalid ID")
	}
	return db.Save(categoriaNegocio).Error
}

// DeleteCategoriaNegocio elimina una relación existente entre una categoría y un negocio
func DeleteCategoriaNegocio(db *gorm.DB, id uint) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return db.Delete(&CategoriaNegocio{}, id).Error
}
