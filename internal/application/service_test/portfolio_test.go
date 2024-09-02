package servicetest

import (
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/leonzag/treport/internal/application/config"
	"github.com/leonzag/treport/internal/application/service"
	"github.com/leonzag/treport/internal/infrastructure/service/tinvest"
	"github.com/leonzag/treport/pkg/logger/zap"
)

var token string

func init() {
	flag.StringVar(&token, "token", "", "tinvest token")
}

func TestTinvestService(t *testing.T) {
	ctx := context.TODO()

	logger, err := zap.NewLoggerDevelop()
	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Skip("skip: token not specified (-token=\"TINVEST_TOKEN\" required)\n")
	}

	cfg := tinvest.NewConfig(config.AppName, "")

	tinvestService := tinvest.NewTinvestService(logger, cfg)
	portfolioService := service.NewPortfolioService(tinvestService)

	if err := portfolioService.Ping(ctx, token); err != nil {
		t.Fatal(err)
	}

	sums, err := portfolioService.SummaryAll(ctx, token)
	if err != nil {
		t.Fatal(err)
	}
	for i, sum := range sums {
		instrs := ""
		for _, ins := range sum.Instruments {
			instrs += fmt.Sprintf("\t%s\n", ins.Name)
		}
		if instrs == "" {
			instrs = "\tNo instruments"
		}

		info := fmt.Sprintf("Portfolio #%d:\n", i+1)
		info += fmt.Sprintf("Account:\n%s\n", sum.Account)
		info += fmt.Sprintf("Instruments:\n%s\n", instrs)
		info += fmt.Sprintf("Expected Yild: %s\n", sum.Portfolio.ExpectedYield)

		info += fmt.Sprintf("\nSummary created: %s", sum.CreatedAt())

		logger.Infof(info)
	}
}
