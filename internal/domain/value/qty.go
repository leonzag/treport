package value

import (
	"fmt"
	"math"
)

// Quotation Котировка - денежная сумма без указания валюты
type Quotation struct {
	Units int64 // Целая часть суммы, может быть отрицательным числом
	Nano  int32 // Дробная часть суммы, может быть отрицательным числом
}

// ToFloat - get value as float64 number
func (q *Quotation) ToFloat() float64 {
	if q != nil {
		num := float64(q.Units) + float64(q.Nano)*math.Pow10(-9)
		num = num * math.Pow10(9)
		num = math.Round(num)
		num = num / math.Pow10(9)
		return num
	}
	return float64(0)
}

func (q *Quotation) String() string {
	return fmt.Sprint(q.ToFloat())
}
