package enum

type Currency int32

const (
	_ Currency = iota

	CurrencyRUB
	CurrencyUSD
	CurrencyEUR
)

func (c Currency) String() string {
	return [...]string{
		"RUB",
		"USD",
		"EUR",
	}[c-1]
}
