package config

import (
	"github.com/Israel-Ferreira/livros-grpc/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDbConfig() {
	db, err := gorm.Open(sqlite.Open("livros.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.LivroModel{})

	DB = db
}
