package main

import (
	"log"
	"net"

	"github.com/paypay3/go-graphql-grpc-sample/article/repository"

	"google.golang.org/grpc"

	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
	"github.com/paypay3/go-graphql-grpc-sample/article/service"
)

func main() {

	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// Repositoryを作成
	r, err := repository.NewMySQLHandler()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	// Serviceを作成
	s := service.NewService(r)

	//サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, s)

	//articleサーバーを起動
	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
