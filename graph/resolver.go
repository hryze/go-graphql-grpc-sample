package graph

import "github.com/paypay3/go-graphql-grpc-sample/article/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleClient *client.Client
}
