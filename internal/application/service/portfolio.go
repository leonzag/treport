package service

import (
	"context"
	"errors"
	"sync"

	"github.com/leonzag/treport/internal/application/interfaces"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
)

var _ interfaces.PortfolioService = new(portfolioService)

type portfolioService struct {
	tinvest  interfaces.TinvestAPI
	useCache bool
}

func NewPortfolioService(tinvest interfaces.TinvestAPI) *portfolioService {
	return &portfolioService{
		tinvest: tinvest,
	}
}

func (s *portfolioService) Ping(ctx context.Context, token string) error {
	if s.tinvest.ActiveConnection() && s.tinvest.Token() != token {
		if err := s.tinvest.ClientStop(); err != nil {
			return err
		}
	}

	return s.tinvest.Ping(ctx, token)
}

func (s *portfolioService) SummaryAll(ctx context.Context, token string) ([]*entity.PortfolioSummary, error) {
	if err := s.tinvest.ClientConnection(ctx, token); err != nil {
		return nil, err
	}
	defer s.tinvest.ClientStop()
	s.tinvest.SetUseCache(s.UseCache())

	return s.summaryAll(ctx)
}

func (s *portfolioService) Summary(ctx context.Context, token string, acc *entity.Account) (*entity.PortfolioSummary, error) {
	if err := s.tinvest.ClientConnection(ctx, token); err != nil {
		return nil, err
	}
	defer s.tinvest.ClientStop()
	s.tinvest.SetUseCache(s.UseCache())

	return s.summary(ctx, acc)
}

func (s *portfolioService) summaryAll(ctx context.Context) ([]*entity.PortfolioSummary, error) {
	accs, err := s.tinvest.Accounts(ctx, enum.AccountStatusALL)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var errs error
	portfolios := make([]*entity.PortfolioSummary, len(accs))

	for i, acc := range accs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			portfolio, err := s.summary(ctx, acc)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				portfolios[i] = portfolio
			}
		}()
	}
	wg.Wait()

	if errs != nil {
		return nil, errs
	}

	return portfolios, nil
}

func (s *portfolioService) summary(ctx context.Context, acc *entity.Account) (*entity.PortfolioSummary, error) {
	portfolio, err := s.tinvest.Portfolio(ctx, acc.Id, enum.CurrencyRUB)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var errs error
	instruments := make([]*entity.Instrument, len(portfolio.Positions))

	for i, position := range portfolio.Positions {
		wg.Add(1)

		go func() {
			defer wg.Done()
			instrument, err := s.instrument(ctx, position.InstrumentUid)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				instruments[i] = instrument
			}
		}()
	}
	wg.Wait()

	if errs != nil {
		return nil, errs
	}

	return entity.NewPortfolioSummary(acc, portfolio, instruments), nil
}

func (s *portfolioService) instrument(ctx context.Context, uid string) (*entity.Instrument, error) {
	return s.tinvest.Instrument(ctx, uid)
}

func (s *portfolioService) SetUseCache(use bool) {
	s.useCache = use
}

func (s *portfolioService) UseCache() bool {
	return s.useCache
}
