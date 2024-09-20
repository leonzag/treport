package app

import (
	"context"
	"fmt"
	"time"

	"github.com/leonzag/treport/internal/domain/entity"
)

func (a *application) reportProcess(dest string, token string, timeLimit time.Duration) (string, error) {
	portfolioSrv := a.Services().Portfolio()
	portfolioSrv.SetUseCache(true)
	reportSrv := a.Services().Report()
	sorting := entity.NewPositionsSoringDefault()

	summaryCh := make(chan []*entity.PortfolioSummary)
	errCh := make(chan error)

	ctxWithTimeout, cancel := context.WithTimeout(a.Ctx(), timeLimit)
	defer cancel()

	go func() {
		summary, err := portfolioSrv.SummaryAll(ctxWithTimeout, token)
		if err != nil {
			errCh <- err
			return
		}
		summaryCh <- summary
	}()

	select {
	case <-ctxWithTimeout.Done():
		return "", fmt.Errorf("Не удалось сделать отчет. Медленное соединение.")
	case err := <-errCh:
		return "", err
	case summary := <-summaryCh:
		for _, s := range summary {
			s.Portfolio.SortPositionsByTypes(sorting...)
		}
		return reportSrv.CreateXLSX(dest, summary)
	}
}
