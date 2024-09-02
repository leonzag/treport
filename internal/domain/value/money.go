package value

import "math"

// MoneyValue Денежная сумма в определенной валюте
type MoneyValue struct {
	Currency string
	Units    int64 // Целая часть суммы, может быть отрицательным числом.
	Nano     int32 // Дробная часть суммы, может быть отрицательным числом.
}

// ToFloat - get value as float64 number
func (mv *MoneyValue) ToFloat() float64 {
	if mv != nil {
		num := float64(mv.Units) + float64(mv.Nano)*math.Pow10(-9)
		num = num * math.Pow10(9)
		num = math.Round(num)
		num = num / math.Pow10(9)
		return num
	}
	return float64(0)
}
