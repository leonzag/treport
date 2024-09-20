package app

type appServices struct {
	token     TokenService
	portfolio PortfolioService
	report    PortfolioReportService
}

var _ AppServices = new(appServices)

func NewAppServices(
	token TokenService,
	portfolio PortfolioService,
	report PortfolioReportService,
) *appServices {
	return &appServices{
		token:     token,
		portfolio: portfolio,
		report:    report,
	}
}

func (s *appServices) Token() TokenService {
	return s.token
}

func (s *appServices) Portfolio() PortfolioService {
	return s.portfolio
}

func (s *appServices) Report() PortfolioReportService {
	return s.report
}
