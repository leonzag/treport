package mapper

import (
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type InstrumentMapper struct {
	ValueMapper
}

func (m InstrumentMapper) InstrumentToDomain(i *investapi.Instrument) *entity.Instrument {
	if i == nil {
		return nil
	}
	return &entity.Instrument{
		First_1MinCandleDate: i.GetFirst_1MinCandleDate().AsTime(),
		First_1DayCandleDate: i.GetFirst_1DayCandleDate().AsTime(),

		MinPriceIncrement: m.QuotationToDomain(i.GetMinPriceIncrement()),
		Kshort:            m.QuotationToDomain(i.GetKshort()),
		Klong:             m.QuotationToDomain(i.GetKlong()),
		Dshort:            m.QuotationToDomain(i.GetDshort()),
		Dlong:             m.QuotationToDomain(i.GetDlong()),
		DshortMin:         m.QuotationToDomain(i.GetDshortMin()),
		DlongMin:          m.QuotationToDomain(i.GetDlongMin()),

		Uid:               i.GetUid(),
		Figi:              i.GetFigi(),
		PositionUid:       i.GetPositionUid(),
		AssetUid:          i.GetAssetUid(),
		Name:              i.GetName(),
		Ticker:            i.GetTicker(),
		ClassCode:         i.GetClassCode(),
		Isin:              i.GetIsin(),
		InstrumentType:    i.GetInstrumentType(),
		Exchange:          i.GetExchange(),
		CountryOfRisk:     i.GetCountryOfRisk(),
		CountryOfRiskName: i.GetCountryOfRiskName(),
		Currency:          i.GetCurrency(),

		Lot: 0,

		SellAvailableFlag:     i.GetSellAvailableFlag(),
		BuyAvailableFlag:      i.GetBuyAvailableFlag(),
		ApiTradeAvailableFlag: i.GetApiTradeAvailableFlag(),
		OtcFlag:               i.GetOtcFlag(),
		BlockedTcaFlag:        i.GetBlockedTcaFlag(),
		ShortEnabledFlag:      i.GetShortEnabledFlag(),
		ForIisFlag:            i.GetForIisFlag(),
		ForQualInvestorFlag:   i.GetForQualInvestorFlag(),
		WeekendFlag:           i.GetWeekendFlag(),

		TradingStatus:  m.SecurityTradingStatusToDomain(i.GetTradingStatus()),
		RealExchange:   m.RealExchangeToDomain(i.GetRealExchange()),
		InstrumentKind: m.InstrumentTypeToDomain(i.GetInstrumentKind()),
		// Brand                 *BrandData
	}
}

func (m InstrumentMapper) BrandDataToDomain(b *investapi.BrandData) *entity.BrandData {
	if b == nil {
		return nil
	}
	return &entity.BrandData{
		LogoName:      b.GetLogoName(),
		LogoBaseColor: b.GetLogoBaseColor(),
		TextColor:     b.GetTextColor(),
	}
}

func (m InstrumentMapper) SecurityTradingStatusToDomain(s investapi.SecurityTradingStatus) enum.SecurityTradingStatus {
	switch s {
	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING:
		return enum.SecurityTradingStatus_DEALER_NOT_AVAILABLE_FOR_TRADING

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_PERIOD:
		return enum.SecurityTradingStatus_OPENING_PERIOD

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_PERIOD:
		return enum.SecurityTradingStatus_CLOSING_PERIOD

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_BREAK_IN_TRADING:
		return enum.SecurityTradingStatus_BREAK_IN_TRADING

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_NORMAL_TRADING:
		return enum.SecurityTradingStatus_NORMAL_TRADING

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_AUCTION:
		return enum.SecurityTradingStatus_CLOSING_AUCTION

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DARK_POOL_AUCTION:
		return enum.SecurityTradingStatus_DARK_POOL_AUCTION

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DISCRETE_AUCTION:
		return enum.SecurityTradingStatus_DISCRETE_AUCTION

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD:
		return enum.SecurityTradingStatus_OPENING_AUCTION_PERIOD

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE:
		return enum.SecurityTradingStatus_TRADING_AT_CLOSING_AUCTION_PRICE

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_ASSIGNED:
		return enum.SecurityTradingStatus_SESSION_ASSIGNED

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_CLOSE:
		return enum.SecurityTradingStatus_SESSION_CLOSE

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_OPEN:
		return enum.SecurityTradingStatus_SESSION_OPEN

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING:
		return enum.SecurityTradingStatus_DEALER_NORMAL_TRADING

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING:
		return enum.SecurityTradingStatus_DEALER_BREAK_IN_TRADING

	case investapi.SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING:
		return enum.SecurityTradingStatus_DEALER_NOT_AVAILABLE_FOR_TRADING

	default:
		return enum.SecurityTradingStatus_UNSPECIFIED
	}
}

func (m InstrumentMapper) RealExchangeToDomain(e investapi.RealExchange) enum.RealExchange {
	switch e {
	case investapi.RealExchange_REAL_EXCHANGE_MOEX:
		return enum.RealExchange_MOEX
	case investapi.RealExchange_REAL_EXCHANGE_RTS:
		return enum.RealExchange_RTS
	case investapi.RealExchange_REAL_EXCHANGE_OTC:
		return enum.RealExchange_OTC
	case investapi.RealExchange_REAL_EXCHANGE_DEALER:
		return enum.RealExchange_DEALER
	default:
		return enum.RealExchange_UNSPECIFIED
	}
}

func (m InstrumentMapper) InstrumentTypeToDomain(t investapi.InstrumentType) enum.InstrumentType {
	switch t {
	case investapi.InstrumentType_INSTRUMENT_TYPE_BOND:
		return enum.InstrumentType_BOND
	case investapi.InstrumentType_INSTRUMENT_TYPE_SHARE:
		return enum.InstrumentType_SHARE
	case investapi.InstrumentType_INSTRUMENT_TYPE_CURRENCY:
		return enum.InstrumentType_CURRENCY
	case investapi.InstrumentType_INSTRUMENT_TYPE_ETF:
		return enum.InstrumentType_ETF
	case investapi.InstrumentType_INSTRUMENT_TYPE_FUTURES:
		return enum.InstrumentType_FUTURES
	case investapi.InstrumentType_INSTRUMENT_TYPE_SP:
		return enum.InstrumentType_SP
	case investapi.InstrumentType_INSTRUMENT_TYPE_OPTION:
		return enum.InstrumentType_OPTION
	case investapi.InstrumentType_INSTRUMENT_TYPE_CLEARING_CERTIFICATE:
		return enum.InstrumentType_CLEARING_CERTIFICATE
	case investapi.InstrumentType_INSTRUMENT_TYPE_INDEX:
		return enum.InstrumentType_INDEX
	case investapi.InstrumentType_INSTRUMENT_TYPE_COMMODITY:
		return enum.InstrumentType_COMMODITY
	default:
		return enum.InstrumentType_UNSPECIFIED
	}
}
