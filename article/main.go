package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/paypay3/go-graphql-grpc-sample/article/client"
	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
)

// 動作確認用
func main() {
	// clientを生成
	c, _ := client.NewClient("localhost:50051")

	//createArticle(c)
	//readArticle(c)
	//updateArticle(c)
	deleteArticle(c)
	listArticles(c)
}

func createArticle(c *client.Client) {
	input := &pb.ArticleInput{
		Author:  "gopher",
		Title:   "gRPC",
		Content: "gRPC is so cool!",
	}

	res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to CreateArticle: %v", err)
	}

	fmt.Printf("CreateArticle Response: \n%v\n", res)
}

func readArticle(c *client.Client) {
	var id int64 = 1
	res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("Failed to ReadArticle: %v", err)
	}

	fmt.Printf("ReadArticle Response: \n%v\n", res)
}

func updateArticle(c *client.Client) {
	var id int64 = 5
	input := &pb.ArticleInput{
		Author:  "GraphQL master",
		Title:   "GraphQL",
		Content: "GraphQL is very smart!",
	}

	res, err := c.Service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to UpdateArticle: %v", err)
	}

	fmt.Printf("UpdateArticle Response: \n%v\n", res)
}

func deleteArticle(c *client.Client) {
	var id int64 = 4
	res, err := c.Service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("Failed to UpdateArticle: %v", err)
	}

	fmt.Printf("The article has been deleted \n%v\n", res)
}

func listArticles(c *client.Client) {
	stream, err := c.Service.ListArticle(context.Background(), &pb.ListArticleRequest{})
	if err != nil {
		log.Fatalf("Failed to ListArticle: %v", err)
	}

	// Server Streamingで渡されたレスポンスを１つ１つ受け取る
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to Server Streaming: %v", err)
		}

		fmt.Println(res)
	}
}
