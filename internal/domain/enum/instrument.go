package enum

// Режим торгов инструмента
type (
	SecurityTradingStatus int32
	RealExchange          int32
	InstrumentType        int32
)

const (
	_ SecurityTradingStatus = iota

	SecurityTradingStatus_UNSPECIFIED                      // Торговый статус не определён.
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING        // Недоступен для торгов.
	SecurityTradingStatus_OPENING_PERIOD                   // Период открытия торгов.
	SecurityTradingStatus_CLOSING_PERIOD                   // Период закрытия торгов.
	SecurityTradingStatus_BREAK_IN_TRADING                 // Перерыв в торговле.
	SecurityTradingStatus_NORMAL_TRADING                   // Нормальная торговля.
	SecurityTradingStatus_CLOSING_AUCTION                  // Аукцион закрытия.
	SecurityTradingStatus_DARK_POOL_AUCTION                // Аукцион крупных пакетов.
	SecurityTradingStatus_DISCRETE_AUCTION                 // Дискретный аукцион.
	SecurityTradingStatus_OPENING_AUCTION_PERIOD           // Аукцион открытия.
	SecurityTradingStatus_TRADING_AT_CLOSING_AUCTION_PRICE // Период торгов по цене аукциона закрытия.
	SecurityTradingStatus_SESSION_ASSIGNED                 // Сессия назначена.
	SecurityTradingStatus_SESSION_CLOSE                    // Сессия закрыта.
	SecurityTradingStatus_SESSION_OPEN                     // Сессия открыта.
	SecurityTradingStatus_DEALER_NORMAL_TRADING            // Доступна торговля в режиме внутренней ликвидности брокера.
	SecurityTradingStatus_DEALER_BREAK_IN_TRADING          // Перерыв торговли в режиме внутренней ликвидности брокера.
	SecurityTradingStatus_DEALER_NOT_AVAILABLE_FOR_TRADING // Недоступна торговля в режиме внутренней ликвидности брокера.
)

func (s SecurityTradingStatus) String() string {
	return [...]string{
		"SECURITY_TRADING_STATUS_UNSPECIFIED",
		"SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING",
		"SECURITY_TRADING_STATUS_OPENING_PERIOD",
		"SECURITY_TRADING_STATUS_CLOSING_PERIOD",
		"SECURITY_TRADING_STATUS_BREAK_IN_TRADING",
		"SECURITY_TRADING_STATUS_NORMAL_TRADING",
		"SECURITY_TRADING_STATUS_CLOSING_AUCTION",
		"SECURITY_TRADING_STATUS_DARK_POOL_AUCTION",
		"SECURITY_TRADING_STATUS_DISCRETE_AUCTION",
		"SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD",
		"SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE",
		"SECURITY_TRADING_STATUS_SESSION_ASSIGNED",
		"SECURITY_TRADING_STATUS_SESSION_CLOSE",
		"SECURITY_TRADING_STATUS_SESSION_OPEN",
		"SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING",
		"SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING",
		"SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING",
	}[s-1]
}

const (
	_ RealExchange = iota

	RealExchange_UNSPECIFIED // Тип не определён.
	RealExchange_MOEX        // Московская биржа.
	RealExchange_RTS         // Санкт-Петербургская биржа.
	RealExchange_OTC         // Внебиржевой инструмент.
	RealExchange_DEALER      // Инструмент, торгуемый на площадке брокера.
)

func (e RealExchange) String() string {
	return [...]string{
		"REAL_EXCHANGE_UNSPECIFIED",
		"REAL_EXCHANGE_MOEX",
		"REAL_EXCHANGE_RTS",
		"REAL_EXCHANGE_OTC",
		"REAL_EXCHANGE_DEALER",
	}[e-1]
}

const (
	_ InstrumentType = iota

	InstrumentType_UNSPECIFIED
	InstrumentType_BOND                 // Облигация.
	InstrumentType_SHARE                // Акция.
	InstrumentType_CURRENCY             // Валюта.
	InstrumentType_ETF                  // Exchange-traded fund. Фонд.
	InstrumentType_FUTURES              // Фьючерс.
	InstrumentType_SP                   // Структурная нота.
	InstrumentType_OPTION               // Опцион.
	InstrumentType_CLEARING_CERTIFICATE // Clearing certificate.
	InstrumentType_INDEX                // Индекс.
	InstrumentType_COMMODITY            // Товар.
)

func (t InstrumentType) String() string {
	return [...]string{
		"INSTRUMENT_TYPE_UNSPECIFIED",
		"INSTRUMENT_TYPE_BOND",
		"INSTRUMENT_TYPE_SHARE",
		"INSTRUMENT_TYPE_CURRENCY",
		"INSTRUMENT_TYPE_ETF",
		"INSTRUMENT_TYPE_FUTURES",
		"INSTRUMENT_TYPE_SP",
		"INSTRUMENT_TYPE_OPTION",
		"INSTRUMENT_TYPE_CLEARING_CERTIFICATE",
		"INSTRUMENT_TYPE_INDEX",
		"INSTRUMENT_TYPE_COMMODITY",
	}[t-1]
}
