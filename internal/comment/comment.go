package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("error while fetching comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Printf("fetching comments\n")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Printf("error in getting comment\n")
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
