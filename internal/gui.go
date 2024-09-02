package internal

import (
	"context"
	"database/sql"

	"github.com/leonzag/treport/internal/application/config"
	appInterface "github.com/leonzag/treport/internal/application/interfaces"
	appService "github.com/leonzag/treport/internal/application/service"
	tokenRepo "github.com/leonzag/treport/internal/infrastructure/repo/token/sqlite"
	"github.com/leonzag/treport/internal/infrastructure/service/tinvest"
	guiApp "github.com/leonzag/treport/internal/presentation/gui/app"
	guiInterface "github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/pkg/database/sqlite"
	"github.com/leonzag/treport/pkg/logger"
	"github.com/leonzag/treport/pkg/logger/zerolog"
)

type appGuiDeps struct {
	ctx    context.Context
	logger logger.Logger
	db     *sql.DB

	tokenRepo              appInterface.TokenRepo
	tokenService           appInterface.TokenService
	cryptoService          appInterface.CryptoService
	tinvestService         appInterface.TinvestAPI
	portfolioService       appInterface.PortfolioService
	portfolioReportService appInterface.PortfolioReportService

	appServices guiInterface.AppServices
}

type AppGui interface {
	ShowAndRun() error
}

func NewAppGUI() (AppGui, error) {
	deps := &appGuiDeps{}

	ctx := deps.Ctx()
	logger, err := deps.Logger()
	if err != nil {
		return nil, err
	}
	services, err := deps.AppServices()
	if err != nil {
		return nil, err
	}
	app := guiApp.NewApp(ctx, logger, services)

	return app, nil
}

func (a *appGuiDeps) Ctx() context.Context {
	if a.ctx != nil {
		return a.ctx
	}
	a.ctx = context.Background()

	return a.ctx
}

func (a *appGuiDeps) Logger() (logger.Logger, error) {
	if a.logger != nil {
		return a.logger, nil
	}
	logger, err := zerolog.NewLogger()
	if err != nil {
		return nil, err
	}
	a.logger = logger

	return logger, nil
}

func (a *appGuiDeps) Db() (*sql.DB, error) {
	if a.db != nil {
		return a.db, nil
	}
	dbPath, err := config.SQLiteDBPathDefault(config.AppName)
	if err != nil {
		return nil, err
	}
	db, err := sqlite.New(dbPath)
	if err != nil {
		return nil, err
	}
	a.db = db

	return db, nil
}

func (a *appGuiDeps) TokenRepo() (appInterface.TokenRepo, error) {
	if a.tokenRepo != nil {
		return a.tokenRepo, nil
	}
	db, err := a.Db()
	if err != nil {
		return nil, err
	}
	repo := tokenRepo.NewTokenRepo(db)
	if !repo.IsInited() {
		if err := repo.Init(a.Ctx()); err != nil {
			return nil, err
		}
	}
	a.tokenRepo = repo

	return repo, nil
}

func (a *appGuiDeps) TokenService() (appInterface.TokenService, error) {
	if a.tokenService != nil {
		return a.tokenService, nil
	}
	token, err := a.TokenRepo()
	if err != nil {
		return nil, err
	}
	crypto, err := a.CryptoService()
	if err != nil {
		return nil, err
	}
	a.tokenService = appService.NewTokenService(token, crypto)

	return a.tokenService, nil
}

func (a *appGuiDeps) CryptoService() (appInterface.CryptoService, error) {
	if a.cryptoService != nil {
		return a.cryptoService, nil
	}
	a.cryptoService = appService.NewCryptoService()

	return a.cryptoService, nil
}

func (a *appGuiDeps) TinvestService() (appInterface.TinvestAPI, error) {
	if a.tinvestService != nil {
		return a.tinvestService, nil
	}
	l, err := a.Logger()
	if err != nil {
		return nil, err
	}
	a.tinvestService = tinvest.NewTinvestService(l, tinvest.NewConfig(config.AppName, ""))

	return a.tinvestService, nil
}

func (a *appGuiDeps) PortfolioReportService() (appInterface.PortfolioReportService, error) {
	if a.portfolioReportService != nil {
		return a.portfolioReportService, nil
	}
	l, err := a.Logger()
	if err != nil {
		return nil, err
	}
	a.portfolioReportService = appService.NewPortfolioReportService(l)

	return a.portfolioReportService, nil
}

func (a *appGuiDeps) PortfolioService() (appInterface.PortfolioService, error) {
	if a.portfolioService != nil {
		return a.portfolioService, nil
	}
	t, err := a.TinvestService()
	if err != nil {
		return nil, err
	}
	a.portfolioService = appService.NewPortfolioService(t)

	return a.portfolioService, nil
}

func (a *appGuiDeps) AppServices() (guiInterface.AppServices, error) {
	if a.appServices != nil {
		return a.appServices, nil
	}

	token, err := a.TokenService()
	if err != nil {
		return nil, err
	}
	portfolio, err := a.PortfolioService()
	if err != nil {
		return nil, err
	}
	report, err := a.PortfolioReportService()
	if err != nil {
		return nil, err
	}
	a.appServices = guiApp.NewAppServices(token, portfolio, report)

	return a.appServices, nil
}
