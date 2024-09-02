package interfaces

import (
	"context"

	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
)

type CryptoService interface {
	HashPassword(pwd string) (string, error)
	CheckPassword(hashed string, pwd string) bool
	EncryptToken(pwd string, token string) (string, error)
	DecryptToken(pwd string, encryptedToken string) (string, error)
}

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

type TinvestAPI interface {
	ClientConnection(ctx context.Context, token string) error
	ClientStop() error
	Ping(ctx context.Context, token string) error
	Accounts(ctx context.Context, status enum.AccountStatus) ([]*entity.Account, error)
	Instrument(ctx context.Context, uid string) (*entity.Instrument, error)
	Portfolio(ctx context.Context, accId string, crc enum.Currency) (*entity.Portfolio, error)
	ActiveConnection() bool
	SetUseCache(use bool)
	UseCache() bool
	Token() string
}
