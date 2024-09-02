package app

import "github.com/leonzag/treport/internal/presentation/gui/interfaces"

type appServices struct {
	token     interfaces.TokenService
	portfolio interfaces.PortfolioService
	report    interfaces.PortfolioReportService
}

var _ interfaces.AppServices = new(appServices)

func NewAppServices(
	token interfaces.TokenService,
	portfolio interfaces.PortfolioService,
	report interfaces.PortfolioReportService,
) *appServices {
	return &appServices{
		token:     token,
		portfolio: portfolio,
		report:    report,
	}
}

func (s *appServices) Token() interfaces.TokenService {
	return s.token
}

func (s *appServices) Portfolio() interfaces.PortfolioService {
	return s.portfolio
}

func (s *appServices) Report() interfaces.PortfolioReportService {
	return s.report
}
