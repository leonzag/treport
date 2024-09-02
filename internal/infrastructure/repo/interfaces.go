package repo

import (
	"context"

	"github.com/leonzag/treport/internal/domain/entity"
)

type TokenRepo interface {
	Init(ctx context.Context) error
	IsInited() bool
	Close() error

	Add(ctx context.Context, token *entity.Token) error
	Get(ctx context.Context, name string) (*entity.Token, error)
	List(ctx context.Context) ([]*entity.Token, error)
	Update(ctx context.Context, token *entity.Token) error
	Delete(ctx context.Context, token *entity.Token) error
}
