package servicetest

import (
	"time"

	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/leonzag/treport/internal/domain/value"
)

func newPortfolioSummaryExample() *entity.PortfolioSummary {
	openedDate, _ := time.Parse(time.RFC1123Z, time.RFC1123Z)
	accId := "1000"
	a := &entity.Account{
		OpenedDate:  openedDate,
		Id:          accId,
		Name:        "Пример Счета",
		Type:        enum.AccountTypeTINKOFF,
		Status:      enum.AccountStatusOPEN,
		AccessLevel: enum.AccessLevelFULLACCESS,
	}
	p := &entity.Portfolio{
		TotalAmountShares:    &value.MoneyValue{Currency: "RUB", Units: 10000, Nano: 0},
		TotalAmountPortfolio: &value.MoneyValue{Currency: "RUB", Units: 10000, Nano: 0},
		ExpectedYield:        &value.Quotation{Units: 100, Nano: 0},

		AccountId: accId,
	}
	is := []*entity.Instrument{}
	return entity.NewPortfolioSummary(a, p, is)
}
