package entity

import (
	"time"

	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/leonzag/treport/internal/domain/value"
)

type Instrument struct {
	First_1MinCandleDate time.Time
	First_1DayCandleDate time.Time

	MinPriceIncrement *value.Quotation
	Kshort            *value.Quotation
	Klong             *value.Quotation
	Dshort            *value.Quotation
	Dlong             *value.Quotation
	DshortMin         *value.Quotation
	DlongMin          *value.Quotation
	Brand             *BrandData

	Uid               string
	Figi              string
	PositionUid       string
	AssetUid          string
	Name              string
	Ticker            string
	ClassCode         string
	Isin              string
	InstrumentType    string
	Exchange          string
	CountryOfRisk     string
	CountryOfRiskName string
	Currency          string

	Lot int32

	SellAvailableFlag     bool
	BuyAvailableFlag      bool
	ApiTradeAvailableFlag bool
	OtcFlag               bool
	BlockedTcaFlag        bool
	ShortEnabledFlag      bool
	ForIisFlag            bool
	ForQualInvestorFlag   bool
	WeekendFlag           bool

	TradingStatus  enum.SecurityTradingStatus
	RealExchange   enum.RealExchange
	InstrumentKind enum.InstrumentType
}

type BrandData struct {
	LogoName      string
	LogoBaseColor string
	TextColor     string
}
