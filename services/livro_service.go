package services

import (
	"context"

	"github.com/Israel-Ferreira/livros-grpc/models"
	"github.com/Israel-Ferreira/livros-grpc/repositories"
	"github.com/Israel-Ferreira/livros-grpc/services/pb"
)

func convertLivroModelToLivro(livro models.LivroModel) *pb.Livro {
	return &pb.Livro{
		Id:     int64(livro.ID),
		Titulo: livro.Titulo,
		Isbn:   livro.ISBN,
		Autor:  livro.Autor,
		Preco:  livro.Preco,
	}
}

type LivroServer struct {
	lvrRepo repositories.LivroRepository
	pb.UnimplementedLivroServiceServer
}

func (lvs *LivroServer) FindById(ctx context.Context, idReq *pb.FindByIdRequest) (*pb.Livro, error) {
	livro, err := lvs.lvrRepo.FindById(uint64(idReq.Id))

	if err != nil {
		return nil, err
	}

	livroResp := convertLivroModelToLivro(livro)

	return livroResp, nil
}

func (lvs *LivroServer) FindAll(ctx context.Context, emptyReq *pb.EmptyRequest) (*pb.LivroList, error) {
	livros, err := lvs.lvrRepo.FindAll()

	if err != nil {
		return nil, err
	}

	var livrosList []*pb.Livro

	for _, livro := range livros {
		livroResp := convertLivroModelToLivro(livro)
		livrosList = append(livrosList, livroResp)
	}

	return &pb.LivroList{
		Livros: livrosList,
	}, nil
}

func (lvs *LivroServer) DeleteById(ctx context.Context, idReq *pb.FindByIdRequest) (*pb.Livro, error) {
	livro, err := lvs.lvrRepo.DeleteById(uint64(idReq.Id))

	if err != nil {
		return nil, err
	}

	pbLivro := convertLivroModelToLivro(livro)

	return pbLivro, nil
}

func (lvs *LivroServer) Create(ctx context.Context, livroReq *pb.LivroRequest) (*pb.Livro, error) {
	livro, err := lvs.lvrRepo.SaveLivro(livroReq)

	if err != nil {
		return nil, err
	}

	livroResp := convertLivroModelToLivro(livro)
	return livroResp, nil
}

func NewLivroServer(repo repositories.LivroRepository) LivroServer {
	return LivroServer{
		lvrRepo: repo,
	}
}
