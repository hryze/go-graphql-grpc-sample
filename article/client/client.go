package client

import (
	"google.golang.org/grpc"

	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
)

type Client struct {
	conn    *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewArticleServiceClient(conn)

	return &Client{conn, c}, nil
}
