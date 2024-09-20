package app

import (
	"context"

	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/pkg/logger"
)

type AppServices interface {
	Token() TokenService
	Portfolio() PortfolioService
	Report() PortfolioReportService
}

type App interface {
	Ctx() context.Context
	Services() AppServices
	Logger() logger.Logger

	ShowAndRun() error
	Refresh() error
	Quit()

	AddToken(token dto.TokenRequestDTO)
	DeleteToken(token dto.TokenRequestDTO)
	CreateReport(token dto.TokenRequestDTO)
	ToScreenAddToken()
	ToScreenCreateReport()
}

type TokenService interface {
	AddToken(ctx context.Context, token dto.TokenRequestDTO) (*entity.Token, error)

	GetTokenByTitle(ctx context.Context, title string) (*entity.Token, error)
	DecryptToken(token *entity.Token, pwd string) (*entity.TokenDecrypted, error)
	ListTokens(ctx context.Context) ([]*entity.Token, error)
	ListTokensTitles(ctx context.Context) ([]string, error)

	UpdateToken(ctx context.Context, token dto.TokenRequestDTO) (*entity.Token, error)
	DeleteToken(ctx context.Context, token dto.TokenRequestDTO) error
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
