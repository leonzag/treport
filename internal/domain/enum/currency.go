package enum

type Currency int32

const (
	CurrencyRUB Currency = iota
	CurrencyUSD
	CurrencyEUR
)

func (c Currency) String() string {
	return [...]string{
		"RUB",
		"USD",
		"EUR",
	}[c]
}
