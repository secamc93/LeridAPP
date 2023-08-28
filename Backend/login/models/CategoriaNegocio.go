package models

import (
	"gorm.io/gorm"
)

type CategoriaNegocio struct {
	gorm.Model
	NegocioID   uint
	CategoríaID uint
	Negocio     Negocio   `gorm:"foreignkey:NegocioID"`
	Categoría   Categoria `gorm:"foreignkey:CategoríaID"`
}
