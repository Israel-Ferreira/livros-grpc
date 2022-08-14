package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Israel-Ferreira/livros-grpc/config"
	"github.com/Israel-Ferreira/livros-grpc/repositories"
	"github.com/Israel-Ferreira/livros-grpc/services"
	"github.com/Israel-Ferreira/livros-grpc/services/pb"
	"google.golang.org/grpc"
)

func init() {
	config.LoadDbConfig()
}

func main() {

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalln("Erro ao iniciar o servidor")
	}

	grpcServer := grpc.NewServer()

	fmt.Println("Hello World")

	repo := repositories.LivroRepository{Db: config.DB}

	livroServer := services.NewLivroServer(repo)

	pb.RegisterLivroServiceServer(grpcServer, &livroServer)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
