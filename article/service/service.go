package service

import (
	"context"

	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
	"github.com/paypay3/go-graphql-grpc-sample/article/repository"
)

type Service interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error)
	ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error)
	ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error
}

type service struct {
	repository repository.Repository
	pb.UnimplementedArticleServiceServer
}

func NewService(r repository.Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	input := req.GetArticleInput()

	id, err := s.repository.InsertArticle(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	id := req.GetId()

	a, err := s.repository.SelectArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	}, nil
}

func (s *service) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	id := req.GetId()

	input := req.GetArticleInput()

	if err := s.repository.UpdateArticle(ctx, id, input); err != nil {
		return nil, err
	}

	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	id := req.GetId()

	if err := s.repository.DeleteArticle(ctx, id); err != nil {
		return nil, err
	}

	return &pb.DeleteArticleResponse{
		Id: id,
	}, nil
}

func (s *service) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	rows, err := s.repository.SelectAllArticles()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var article pb.Article
		err := rows.StructScan(&article)
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.ListArticleResponse{Article: &article}); err != nil {
			return err
		}
	}

	return nil
}
