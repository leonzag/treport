package service

import (
	"github.com/leonzag/treport/internal/application/interfaces"
	"github.com/leonzag/treport/internal/application/service/report"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/pkg/logger"
)

var _ interfaces.PortfolioReportService = new(portfolioReportService)

type portfolioReportService struct {
	logger logger.Logger
}

func NewPortfolioReportService(l logger.Logger) *portfolioReportService {
	return &portfolioReportService{
		logger: l,
	}
}

func (s *portfolioReportService) CreateXLSX(fpath string, portfolios []*entity.PortfolioSummary) (string, error) {
	ctl := report.NewPortfolioRerpotController(s.logger)

	return ctl.Report(fpath, portfolios)
}
