package entity

import (
	"time"

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
