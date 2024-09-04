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
		"Торговый статус не определён",
		"Недоступен для торгов",
		"Период открытия торгов",
		"Период закрытия торгов",
		"Перерыв в торговле",
		"Нормальная торговля",
		"Аукцион закрытия",
		"Аукцион крупных пакетов",
		"Дискретный аукцион",
		"Аукцион открытия",
		"Период торгов по цене аукциона закрытия",
		"Сессия назначена",
		"Сессия закрыта",
		"Сессия открыта",
		"Доступна торговля в режиме внутренней ликвидности брокера",
		"Перерыв торговли в режиме внутренней ликвидности брокера",
		"Недоступна торговля в режиме внутренней ликвидности брокера",
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
		"Тип не определён",
		"Московская биржа",
		"Санкт-Петербургская биржа",
		"Внебиржевой инструмент",
		"Инструмент, торгуемый на площадке брокера",
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

func (t InstrumentType) StringRU() string {
	return [...]string{
		"Тип инструмента не определён",
		"Облигация",
		"Акция",
		"Валюта",
		"ETF",
		"Фьючерс",
		"Структурная нота",
		"Опцион",
		"Clearing certificate",
		"Индекс",
		"Товар",
	}[t-1]
}

func (t InstrumentType) String() string {
	return [...]string{
		"unspecified",
		"bond",
		"share",
		"currency",
		"etf",
		"futures",
		"sp",
		"option",
		"clearing_certificate",
		"index",
		"commodity",
	}[t-1]
}

func InstrumentTypeFromString(s string) InstrumentType {
	it, ok := map[string]InstrumentType{
		"unspecified":          InstrumentType_UNSPECIFIED,
		"bond":                 InstrumentType_BOND,
		"share":                InstrumentType_SHARE,
		"currency":             InstrumentType_CURRENCY,
		"etf":                  InstrumentType_ETF,
		"futures":              InstrumentType_FUTURES,
		"sp":                   InstrumentType_SP,
		"option":               InstrumentType_OPTION,
		"clearing_certificate": InstrumentType_CLEARING_CERTIFICATE,
		"index":                InstrumentType_INDEX,
		"commodity":            InstrumentType_COMMODITY,
	}[s]
	if !ok {
		return InstrumentType_UNSPECIFIED
	}
	return it
}
