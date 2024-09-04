package entitytest

import (
	"slices"
	"testing"

	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
)

func TestPortfolioSort(t *testing.T) {
	typesWantOrder := []string{
		enum.InstrumentType_SHARE.String(),
		enum.InstrumentType_SHARE.String(),
		enum.InstrumentType_SHARE.String(),
		enum.InstrumentType_BOND.String(),
		enum.InstrumentType_BOND.String(),
		enum.InstrumentType_CURRENCY.String(),
		enum.InstrumentType_CURRENCY.String(),
		enum.InstrumentType_ETF.String(),
		enum.InstrumentType_UNSPECIFIED.String(),
	}
	portfolio := &entity.Portfolio{
		Positions: []*entity.PortfolioPosition{
			{InstrumentType: enum.InstrumentType_SHARE.String()},
			{InstrumentType: enum.InstrumentType_BOND.String()},
			{InstrumentType: enum.InstrumentType_CURRENCY.String()},
			{InstrumentType: enum.InstrumentType_SHARE.String()},
			{InstrumentType: enum.InstrumentType_CURRENCY.String()},
			{InstrumentType: enum.InstrumentType_BOND.String()},
			{InstrumentType: enum.InstrumentType_ETF.String()},
			{InstrumentType: enum.InstrumentType_UNSPECIFIED.String()},
			{InstrumentType: enum.InstrumentType_SHARE.String()},
		},
	}

	portfolio.SortPositionsByTypes(
		enum.InstrumentType_SHARE,
		enum.InstrumentType_BOND,
		enum.InstrumentType_CURRENCY,
	)

	typesGotOrder := []string{}
	for _, pos := range portfolio.Positions {
		typesGotOrder = append(typesGotOrder, pos.InstrumentType)
	}

	if !slices.Equal(typesGotOrder, typesWantOrder) {
		t.Fatalf("incorrect sorting order: got %v, want %v", typesGotOrder, typesWantOrder)
	}
}
