package repository

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"

	"github.com/paypay3/go-graphql-grpc-sample/article/config"
	"github.com/paypay3/go-graphql-grpc-sample/article/pb"
)

type Repository interface {
	InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error)
	SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error)
	UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error
	DeleteArticle(ctx context.Context, id int64) error
	SelectAllArticles() (*sqlx.Rows, error)
}

type mySQLHandler struct {
	conn *sqlx.DB
}

func NewMySQLHandler() (*mySQLHandler, error) {
	conn, err := sqlx.Open("mysql", config.Env.MySQL.Dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(config.Env.MySQL.MaxConn)
	conn.SetMaxIdleConns(config.Env.MySQL.MaxIdleConn)
	conn.SetConnMaxLifetime(config.Env.MySQL.MaxConnLifetime)

	return &mySQLHandler{
		conn: conn,
	}, nil
}

func (r *mySQLHandler) InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	query := `
        INSERT INTO articles
            (author, title, content)
        VALUES
            (?, ?, ?)`

	result, err := r.conn.Exec(query, input.Author, input.Title, input.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *mySQLHandler) SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error) {
	query := `
        SELECT
            id,
            author,
            title,
            content
        FROM
            articles
        WHERE
            id = ?`

	var article pb.Article
	if err := r.conn.QueryRowx(query, id).StructScan(&article); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &article, nil
}

func (r *mySQLHandler) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error {
	query := `
        UPDATE
            articles
        SET
            author = ?,
            title = ?,
            content = ?
        WHERE 
            id = ?`

	if _, err := r.conn.Exec(query, input.Author, input.Title, input.Content, id); err != nil {
		return err
	}

	return nil
}

func (r *mySQLHandler) DeleteArticle(ctx context.Context, id int64) error {
	query := `
        DELETE
        FROM
            articles
        WHERE
            id = ?`

	if _, err := r.conn.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (r *mySQLHandler) SelectAllArticles() (*sqlx.Rows, error) {
	query := `
        SELECT
            id,
            author,
            title,
            content
        FROM
            articles`

	rows, err := r.conn.Queryx(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
