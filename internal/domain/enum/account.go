package enum

type (
	AccountStatus int32
	AccountType   int32
	AccessLevel   int32
)

const (
	_ AccountStatus = iota

	AccountStatusUNSPECIFIED // Статус счёта не определён.
	AccountStatusNEW         // Новый, в процессе открытия.
	AccountStatusOPEN        // Открытый и активный счёт.
	AccountStatusCLOSED      // Закрытый счёт.
	AccountStatusALL         // Все счета.
)

const (
	_ AccountType = iota

	AccountTypeUNSPECIFIED // Тип аккаунта не определён.
	AccountTypeTINKOFF     // Брокерский счёт Т-Инвестиций.
	AccountTypeTINKOFF_IIS // ИИС.
	AccountTypeINVEST_BOX  // Инвесткопилка.
	AccountTypeINVEST_FUND // Фонд денежного рынка.
)

const (
	_ AccessLevel = iota

	AccessLevelUNSPECIFIED // Уровень доступа не определён.
	AccessLevelFULLACCESS  // Полный доступ к счёту.
	AccessLevelREADONLY    // Доступ с уровнем прав «только чтение».
	AccessLevelNOACCESS    // Доступа нет.
)

func (s AccountStatus) String() string {
	return [...]string{
		"ACCOUNT_STATUS_UNSPECIFIED",
		"ACCOUNT_STATUS_NEW",
		"ACCOUNT_STATUS_OPEN",
		"ACCOUNT_STATUS_CLOSED",
		"ACCOUNT_STATUS_ALL",
	}[s-1]
}

func (s AccountType) String() string {
	return [...]string{
		"ACCOUNT_TYPE_UNSPECIFIED",
		"ACCOUNT_TYPE_TINKOFF",
		"ACCOUNT_TYPE_TINKOFF_IIS",
		"ACCOUNT_TYPE_INVEST_BOX",
		"ACCOUNT_TYPE_INVEST_FUND",
	}[s-1]
}

func (s AccessLevel) String() string {
	return [...]string{
		"ACCOUNT_ACCESS_LEVEL_UNSPECIFIED",
		"ACCOUNT_ACCESS_LEVEL_FULL_ACCESS",
		"ACCOUNT_ACCESS_LEVEL_READ_ONLY",
		"ACCOUNT_ACCESS_LEVEL_NO_ACCESS",
	}[s-1]
}
