package servicetest

import (
	"fmt"
	"testing"

	"github.com/leonzag/treport/internal/application/service/report"
	"github.com/leonzag/treport/internal/domain/entity"
)

type testLogger struct {
	*testing.T
}

func (l testLogger) Debugf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.Log(msg)
}

func (l testLogger) Infof(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.Log(msg)
}

func (l testLogger) Warnf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.Log(msg)
}

func (l testLogger) Errorf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.Log(msg)
}

func (l testLogger) Fatalf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.Log(msg)
}

func TestReportPortfoliosXLSX(t *testing.T) {
	logger := testLogger{t}
	report := report.NewPortfolioRerpotController(logger)

	file, err := report.Report("./", []*entity.PortfolioSummary{newPortfolioSummaryExample()})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("created file: %s", file)
	}
}
