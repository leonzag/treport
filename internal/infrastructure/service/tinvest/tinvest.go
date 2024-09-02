package tinvest

import (
	"context"

	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/leonzag/treport/internal/infrastructure/service"
	"github.com/leonzag/treport/internal/infrastructure/service/tinvest/cache"
	"github.com/leonzag/treport/internal/infrastructure/service/tinvest/mapper"
	"github.com/leonzag/treport/pkg/logger"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
)

var _ service.TinvestAPI = new(tinvestService)

type tinvestService struct {
	conn   *investgo.Client
	cfg    investgo.Config
	logger logger.Logger

	instrumentCache *cache.InstrumentCache
	useCache        bool
	mapper          mapper.TinvestMapper
}

func NewTinvestService(logger logger.Logger, cfg investgo.Config) *tinvestService {
	return &tinvestService{
		logger:          logger,
		cfg:             cfg,
		instrumentCache: cache.NewInstrumentCache(),
	}
}

func (s *tinvestService) ClientConnection(ctx context.Context, token string) error {
	if s.conn != nil && s.Token() != token {
		err := ErrAlreayAuthenticated
		s.logger.Errorf("start tinvest connection failed: %s", err.Error())
		return err
	}
	if s.conn != nil {
		s.logger.Warnf("start tinvest connection: not required (already connected)")
		return nil
	}

	s.cfg.Token = token
	c, err := investgo.NewClient(ctx, s.cfg, s.logger)
	if err != nil {
		s.logger.Errorf("start tinvest connection: %s", err.Error())
		return err
	}
	s.conn = c
	s.logger.Infof("start tinvest connection")

	return nil
}

func (s *tinvestService) ActiveConnection() bool {
	return s.conn != nil
}

func (s *tinvestService) Token() string {
	return s.cfg.Token
}

func (s *tinvestService) ClientStop() error {
	if s.conn == nil {
		s.logger.Warnf("stop tinvest connection: not required")
		return nil
	}

	if err := s.conn.Stop(); err != nil {
		s.logger.Errorf("stop tinvest connection: %s", err.Error())
		return err
	}
	s.conn = nil
	s.logger.Infof("stop tinvest connection: success")

	return nil
}

func (s *tinvestService) UseCache() bool {
	return s.useCache
}

func (s *tinvestService) SetUseCache(use bool) {
	s.logger.Infof("tinvest: set use cache to %t", use)
	s.useCache = use
}

func (s *tinvestService) Ping(ctx context.Context, token string) error {
	s.logger.Infof("try ping check to tinvest")
	prevConn := s.conn

	if err := s.ClientConnection(ctx, token); err != nil {
		s.logger.Errorf("ping check to tinvest failed: %s", err.Error())
		return err
	}
	if prevConn == nil {
		s.logger.Infof("ping check to tinvest: stop after check scheduled")
		defer s.ClientStop()
	}

	if _, err := s.conn.NewUsersServiceClient().GetInfo(); err != nil {
		s.logger.Errorf("ping check to tinvest failed: %s", err.Error())
		return err
	}
	s.logger.Infof("ping check to tinvest success")

	return nil
}

func (s *tinvestService) Instrument(ctx context.Context, uid string) (*entity.Instrument, error) {
	if s.conn == nil {
		if err := s.ClientConnection(ctx, s.cfg.Token); err != nil {
			return nil, err
		}
		defer s.ClientStop()
	}

	return s.instrument(uid)
}

func (s *tinvestService) Accounts(ctx context.Context, status enum.AccountStatus) ([]*entity.Account, error) {
	if s.conn == nil {
		if err := s.ClientConnection(ctx, s.cfg.Token); err != nil {
			return nil, err
		}
		defer s.ClientStop()
	}

	return s.accounts(status)
}

func (s *tinvestService) Portfolio(ctx context.Context, accId string, crc enum.Currency) (*entity.Portfolio, error) {
	if s.conn == nil {

		if err := s.ClientConnection(ctx, s.cfg.Token); err != nil {
			return nil, err
		}
		defer s.ClientStop()
	}

	return s.portfolio(accId, crc)
}

func (s *tinvestService) initCache() {
	if s.instrumentCache == nil {
		s.logger.Debugf("tinvest: init new instrument cache")
		s.instrumentCache = cache.NewInstrumentCache()
	}
}
