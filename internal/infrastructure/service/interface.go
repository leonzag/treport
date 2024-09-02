package service

import (
	"context"

	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
)

type TinvestAPI interface {
	ClientConnection(ctx context.Context, token string) error
	ClientStop() error
	ActiveConnection() bool
	Token() string
	UseCache() bool
	SetUseCache(use bool)

	Ping(ctx context.Context, token string) error
	Accounts(ctx context.Context, status enum.AccountStatus) ([]*entity.Account, error)
	Instrument(ctx context.Context, uid string) (*entity.Instrument, error)
	Portfolio(ctx context.Context, accId string, crc enum.Currency) (*entity.Portfolio, error)
}
