package mapper

import (
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type PortfolioMapper struct {
	ValueMapper
}

func (m PortfolioMapper) CurrencyToRequest(
	c enum.Currency,
) investapi.PortfolioRequest_CurrencyRequest {
	switch c {
	case enum.CurrencyRUB:
		return investapi.PortfolioRequest_RUB

	case enum.CurrencyUSD:
		return investapi.PortfolioRequest_USD

	case enum.CurrencyEUR:
		return investapi.PortfolioRequest_EUR

	default:
		return investapi.PortfolioRequest_RUB
	}
}

func (m PortfolioMapper) CurrencyRequestToDomain(
	c investapi.PortfolioRequest_CurrencyRequest,
) enum.Currency {
	switch c {
	case investapi.PortfolioRequest_RUB:
		return enum.CurrencyRUB

	case investapi.PortfolioRequest_USD:
		return enum.CurrencyUSD

	case investapi.PortfolioRequest_EUR:
		return enum.CurrencyEUR

	default:
		return enum.CurrencyRUB
	}
}

func (m PortfolioMapper) ResponseToDomain(
	p *investgo.PortfolioResponse,
) *entity.Portfolio {
	positions := make([]*entity.PortfolioPosition, len(p.GetPositions()))
	for i, pos := range p.GetPositions() {
		positions[i] = m.PositionToDomain(pos)
	}

	return &entity.Portfolio{
		AccountId:             p.GetAccountId(),
		Positions:             positions,
		ExpectedYield:         m.QuotationToDomain(p.GetExpectedYield()),
		TotalAmountShares:     m.MoneyValueToDomain(p.GetTotalAmountShares()),
		TotalAmountBonds:      m.MoneyValueToDomain(p.GetTotalAmountBonds()),
		TotalAmountEtf:        m.MoneyValueToDomain(p.GetTotalAmountEtf()),
		TotalAmountCurrencies: m.MoneyValueToDomain(p.GetTotalAmountCurrencies()),
		TotalAmountFutures:    m.MoneyValueToDomain(p.GetTotalAmountFutures()),
		TotalAmountOptions:    m.MoneyValueToDomain(p.GetTotalAmountOptions()),
		TotalAmountSp:         m.MoneyValueToDomain(p.GetTotalAmountSp()),
		TotalAmountPortfolio:  m.MoneyValueToDomain(p.GetTotalAmountPortfolio()),
	}
}

func (m PortfolioMapper) PositionToDomain(p *investapi.PortfolioPosition) *entity.PortfolioPosition {
	return &entity.PortfolioPosition{
		Quantity:                 m.QuotationToDomain(p.GetQuantity()),
		AveragePositionPrice:     m.MoneyValueToDomain(p.GetAveragePositionPrice()),
		ExpectedYield:            m.QuotationToDomain(p.GetExpectedYield()),
		CurrentNkd:               m.MoneyValueToDomain(p.GetCurrentNkd()),
		CurrentPrice:             m.MoneyValueToDomain(p.GetCurrentPrice()),
		AveragePositionPriceFifo: m.MoneyValueToDomain(p.GetAveragePositionPriceFifo()),
		BlockedLots:              m.QuotationToDomain(p.GetBlockedLots()),
		VarMargin:                m.MoneyValueToDomain(p.GetVarMargin()),
		ExpectedYieldFifo:        m.QuotationToDomain(p.GetExpectedYieldFifo()),

		// Deprecated: Marked as deprecated in operations.proto.
		AveragePositionPricePt: m.QuotationToDomain(p.GetAveragePositionPricePt()),
		// Deprecated: Marked as deprecated in operations.proto.
		QuantityLots: m.QuotationToDomain(p.GetQuantityLots()),

		Figi:           p.GetFigi(),
		InstrumentType: p.GetInstrumentType(),
		InstrumentUid:  p.GetInstrumentUid(),
		PositionUid:    p.GetPositionUid(),
		Blocked:        p.GetBlocked(),
	}
}

func (m PortfolioMapper) VirtualPositionToDomain(p *investapi.VirtualPortfolioPosition) *entity.VirtualPortfolioPosition {
	return &entity.VirtualPortfolioPosition{
		ExpireDate:               p.GetExpireDate().AsTime(),
		Quantity:                 m.QuotationToDomain(p.GetQuantity()),
		CurrentPrice:             m.MoneyValueToDomain(p.GetCurrentPrice()),
		AveragePositionPrice:     m.MoneyValueToDomain(p.GetAveragePositionPrice()),
		AveragePositionPriceFifo: m.MoneyValueToDomain(p.GetAveragePositionPriceFifo()),
		ExpectedYield:            m.QuotationToDomain(p.GetExpectedYield()),
		ExpectedYieldFifo:        m.QuotationToDomain(p.GetExpectedYieldFifo()),

		Figi:           p.GetFigi(),
		InstrumentType: p.GetInstrumentType(),
		InstrumentUid:  p.GetInstrumentUid(),
		PositionUid:    p.GetPositionUid(),
	}
}
