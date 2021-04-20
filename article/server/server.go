package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
	"github.com/paypay3/go-graphql-grpc-sample/article/repository"
	"github.com/paypay3/go-graphql-grpc-sample/article/service"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	r, err := repository.NewMySQLHandler()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	s := service.NewService(r)

	//サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, s)
	reflection.Register(server)

	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
