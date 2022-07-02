package comment

import (
	"context"
	"fmt"
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
		return Comment{}, err
	}
	return cmt, nil
}
