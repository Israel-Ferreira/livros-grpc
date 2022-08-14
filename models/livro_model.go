package models

import "gorm.io/gorm"

type LivroModel struct {
	gorm.Model
	Titulo string  `gorm:"column:titulo"`
	ISBN   string  `gorm:"column:isbn"`
	Autor  string  `gorm:"autor"`
	Preco  float64 `gorm:"column:preco"`
}
