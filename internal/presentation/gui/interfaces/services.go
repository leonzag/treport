package interfaces

import (
	"context"

	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
)

type TokenService interface {
	AddToken(ctx context.Context, token dto.TokenDTO) error

	GetTokenByTitle(ctx context.Context, title string) (dto.TokenDTO, error)
	GetTokenByTitleDecrypted(ctx context.Context, title string, pwd string) (dto.TokenDTO, error)
	ListTokens(ctx context.Context) ([]dto.TokenDTO, error)
	ListTokensTitles(ctx context.Context) ([]string, error)

	UpdateToken(ctx context.Context, token dto.TokenDTO) error
	DeleteToken(ctx context.Context, token dto.TokenDTO) error
}

type PortfolioService interface {
	Summary(ctx context.Context, token string, acc *entity.Account) (*entity.PortfolioSummary, error)
	SummaryAll(ctx context.Context, token string) ([]*entity.PortfolioSummary, error)
	Ping(ctx context.Context, token string) error
	SetUseCache(bool)
	UseCache() bool
}

type PortfolioReportService interface {
	CreateXLSX(fpath string, portfolios []*entity.PortfolioSummary) (string, error)
}
