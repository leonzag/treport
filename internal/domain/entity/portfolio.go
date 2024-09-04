package entity

import (
	"time"

	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/leonzag/treport/internal/domain/value"
)

type Portfolio struct {
	TotalAmountShares     *value.MoneyValue
	TotalAmountBonds      *value.MoneyValue
	TotalAmountEtf        *value.MoneyValue
	TotalAmountCurrencies *value.MoneyValue
	TotalAmountFutures    *value.MoneyValue
	TotalAmountOptions    *value.MoneyValue
	TotalAmountSp         *value.MoneyValue
	TotalAmountPortfolio  *value.MoneyValue
	ExpectedYield         *value.Quotation

	AccountId        string
	Positions        []*PortfolioPosition
	VirtualPositions []*VirtualPortfolioPosition
}

func (p *Portfolio) SortPositionsByType(t enum.InstrumentType) {
	sorted := make([]*PortfolioPosition, 0, len(p.Positions))
	unsorted := []*PortfolioPosition{}
	for _, pos := range p.Positions {
		instrType := enum.InstrumentTypeFromString(pos.InstrumentType)
		if instrType == t {
			sorted = append(sorted, pos)
		} else {
			unsorted = append(unsorted, pos)
		}
	}
	for _, pos := range unsorted {
		sorted = append(sorted, pos)
	}
	p.Positions = sorted
}

func (p *Portfolio) SortPositionsByTypes(types ...enum.InstrumentType) {
	for i := 0; i < len(types); i++ {
		p.SortPositionsByType(types[len(types)-i-1])
	}
}

type PortfolioPosition struct {
	Quantity             *value.Quotation
	AveragePositionPrice *value.MoneyValue
	ExpectedYield        *value.Quotation
	CurrentNkd           *value.MoneyValue
	CurrentPrice         *value.MoneyValue
	QuantityLots         *value.Quotation
	BlockedLots          *value.Quotation
	VarMargin            *value.MoneyValue
	ExpectedYieldFifo    *value.Quotation

	// Deprecated: Marked as deprecated in operations.proto.
	AveragePositionPriceFifo *value.MoneyValue
	// Deprecated: Marked as deprecated in operations.proto.
	AveragePositionPricePt *value.Quotation

	Figi           string
	InstrumentType string
	InstrumentUid  string
	PositionUid    string
	Blocked        bool
}

type VirtualPortfolioPosition struct {
	ExpireDate               time.Time
	Quantity                 *value.Quotation
	CurrentPrice             *value.MoneyValue
	AveragePositionPrice     *value.MoneyValue
	AveragePositionPriceFifo *value.MoneyValue
	ExpectedYield            *value.Quotation
	ExpectedYieldFifo        *value.Quotation

	Figi           string
	InstrumentType string
	InstrumentUid  string
	PositionUid    string
}
