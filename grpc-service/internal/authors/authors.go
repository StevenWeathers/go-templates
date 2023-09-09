package authors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stevenweathers/go-templates/grpc-service/db"

	authorsv1 "github.com/stevenweathers/go-templates/grpc-service/gen/go/authors/v1"
)

type Service struct {
	authorQueries *db.Queries
	authorsv1.UnimplementedAuthorsServiceServer
}

func (as *Service) createAuthor(ctx context.Context, author *authorsv1.Author) (*authorsv1.Author, error) {
	slog.Info(fmt.Sprintf("rpc request createAuthor(%q)", author))

	bio := pgtype.Text{
		String: author.Bio,
	}

	_, err := as.authorQueries.CreateAuthor(context.Background(), db.CreateAuthorParams{
		Name: author.Name,
		Bio:  bio,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("CreateAuthor failed: %v", err))
		return author, err
	}

	return author, nil
}

func NewServer(authorQueries *db.Queries) authorsv1.AuthorsServiceServer {
	s := &Service{
		authorQueries: authorQueries,
	}

	return s
}
