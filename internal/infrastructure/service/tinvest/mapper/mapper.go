package mapper

import (
	"github.com/leonzag/treport/internal/domain/value"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type TinvestMapper struct {
	ValueMapper
	Account    AccountMapper
	Instrument InstrumentMapper
	Portfolio  PortfolioMapper
}

type ValueMapper struct{}

func (m ValueMapper) MoneyValueToDomain(v *investapi.MoneyValue) *value.MoneyValue {
	return &value.MoneyValue{
		Currency: v.GetCurrency(),
		Units:    v.GetUnits(),
		Nano:     v.GetNano(),
	}
}

func (m ValueMapper) QuotationToDomain(v *investapi.Quotation) *value.Quotation {
	return &value.Quotation{
		Units: v.GetUnits(),
		Nano:  v.GetNano(),
	}
}
