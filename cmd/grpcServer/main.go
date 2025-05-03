package main

import (
	"database/sql"
	"net"

	"github.com.br/leomoritz/gRPC/internal/database"
	"github.com.br/leomoritz/gRPC/internal/pb"
	"github.com.br/leomoritz/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer) // registra o servidor para que o client consiga interagir com o mesmo

	lis, err := net.Listen("tcp", ":50051") // porta padrão do gRPC
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
