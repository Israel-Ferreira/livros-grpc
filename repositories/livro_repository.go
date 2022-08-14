package repositories

import (
	"github.com/Israel-Ferreira/livros-grpc/models"
	"github.com/Israel-Ferreira/livros-grpc/services/pb"
	"gorm.io/gorm"
)

type LivroRepository struct {
	Db *gorm.DB
}

func (lvr LivroRepository) SaveLivro(livroReq *pb.LivroRequest) (models.LivroModel, error) {
	livro := models.LivroModel{
		Titulo: livroReq.Titulo,
		Autor:  livroReq.Autor,
		ISBN:   livroReq.Isbn,
		Preco:  livroReq.Preco,
	}

	if txn := lvr.Db.Save(&livro); txn.Error != nil {
		return models.LivroModel{}, txn.Error
	}

	return livro, nil

}

func (lvr LivroRepository) FindById(id uint64) (models.LivroModel, error) {
	var livro models.LivroModel

	if txn := lvr.Db.First(&livro, "id = ?", id); txn.Error != nil {
		return models.LivroModel{}, txn.Error
	}

	return livro, nil
}

func (lvr LivroRepository) FindAll() ([]models.LivroModel, error) {
	var livros []models.LivroModel

	if txn := lvr.Db.Find(&livros); txn.Error != nil {
		return nil, txn.Error
	}

	return livros, nil
}

func (lvr LivroRepository) DeleteById(id uint64) (models.LivroModel, error) {
	livro, err := lvr.FindById(id)

	if err != nil {
		return models.LivroModel{}, err
	}

	if txn := lvr.Db.Delete(&livro); txn.Error != nil {
		return models.LivroModel{}, txn.Error
	}

	return livro, nil
}
