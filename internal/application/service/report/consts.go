package report

func summaryTitles() []string {
	return []string{
		"Инструмент",       // по Figi/instrument_uid
		"Тикер",            // по Figi/instrument_uid
		"Тип",              // instrument_type: string: Тип инструмента
		"Кол-во (лотов)",   // quantity_lots: Quotation:  Количество лотов в портфеле.
		"Кол-во (шт)",      // quantity: Quotation: Количество инструмента в портфеле в штуках.
		"Ср. цена позиции", // average_position_price: MoneyValue
		"Ср. цена по FIFO", // average_position_price_fifo: MoneyValue

		// Для получения стоимости лота требуется умножить на лотность инструмента.
		"Текущая цена (1 инстр.)",              // current_price: MoneyValue: Текущая цена за 1 инструмент.
		"Текущая доходность",                   // expected_yield: Quotation: Текущая рассчитанная доходность позиции.
		"Текущий НКД",                          // current_nkd: MoneyValue: Текущий НКД.
		"Ср. цена в пунктах (фьюч)",            // average_position_price_pt: Quotation:  Средняя цена позиции в пунктах (фьюч)
		"Текущая доходность (FIFO)",            // expected_yield_fifo: Quotation: Текущая рассчитанная доходность позиции.
		"Блокировано ли? (на бирже)",           // blocked: bool: Заблокировано на бирже.
		"Блокировано (выставленными заявками)", // blocked_lots: Quotation: Кол. бумаг, блок. выставленными заявками.
		"Вариационная маржа",                   // var_margin: MoneyValue: Вариационная маржа
	}
}
