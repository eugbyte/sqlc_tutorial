package service

import (
	"context"

	"github.com/eugbyte/sqlc_tutorial/internal/api/author/repository/codegen"
	"github.com/eugbyte/sqlc_tutorial/internal/domain/model"
	"github.com/eugbyte/sqlc_tutorial/internal/lib"
)

type AuthorRepo = codegen.Querier

type AuthorService struct {
	authorRepo AuthorRepo
}

func NewAuthorService(authorRepo AuthorRepo) *AuthorService {
	return &AuthorService{
		authorRepo: authorRepo,
	}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, arg model.AuthorRequest) (codegen.Author, error) {
	return s.authorRepo.CreateAuthor(ctx, codegen.CreateAuthorParams{
		Name: arg.Name,
		Bio:  lib.ToPgText(arg.Bio),
	})
}

func (s *AuthorService) GetAuthor(ctx context.Context, id int64) (model.AuthorResponse, error) {
	author, err := s.authorRepo.GetAuthor(ctx, id)
	if err != nil {
		return model.AuthorResponse{}, err
	}

	return model.AuthorResponse{
		ID:          author.Author.ID,
		Name:        author.Author.Name,
		Bio:         lib.FromPgText(author.Author.Bio),
		PublisherId: author.Author.PublisherID,
		Publisher: &model.PublisherResponse{
			ID:   author.Publisher.ID,
			Name: author.Publisher.Name,
		},
	}, nil
}

func (s *AuthorService) ListAuthors(ctx context.Context) ([]model.AuthorResponse, error) {
	authors, err := s.authorRepo.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.AuthorResponse, 0, len(authors))
	for _, author := range authors {
		result = append(result, model.AuthorResponse{
			ID:          author.Author.ID,
			Name:        author.Author.Name,
			Bio:         lib.FromPgText(author.Author.Bio),
			PublisherId: author.Author.PublisherID,
			Publisher: &model.PublisherResponse{
				ID:   author.Publisher.ID,
				Name: author.Publisher.Name,
			},
		})
	}

	return result, nil
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, authorId int64, arg model.AuthorRequest) (model.AuthorResponse, error) {
	result, err := s.authorRepo.UpdateAuthor(ctx, codegen.UpdateAuthorParams{
		ID:   authorId,
		Name: arg.Name,
		Bio:  lib.ToPgText(arg.Bio),
	})
	if err != nil {
		return model.AuthorResponse{}, err
	}
	return model.AuthorResponse{
		ID:          result.ID,
		Name:        result.Name,
		PublisherId: result.PublisherID,
		Bio:         lib.FromPgText(result.Bio),
	}, nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id int64) error {
	return s.authorRepo.DeleteAuthor(ctx, id)
}
