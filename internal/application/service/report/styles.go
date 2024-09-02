package report

import (
	"github.com/xuri/excelize/v2"
)

type xlsxStyles struct {
	accName         int
	colName         int
	instrName       int
	currencyGeneral int
	general         int

	yieldLabel         int
	yieldValue         int
	yieldValueNegative int
	yieldValuePositive int

	totalLabel         int
	totalValue         int
	totalValueNegative int
	totalValuePositive int

	borderFull int
}

func newStylesAll(doc *excelize.File) (*xlsxStyles, error) {
	styles := []int{}
	for _, initer := range []func() *excelize.Style{
		newAccNameStyle,
		newColNamesStyle,
		newInstrNameStyle,
		newCurrencyGeneralStyle,
		newGeneralStyle,
		newBorderFull,

		newYieldLabel,
		newYieldValue,
		newYieldValueNegative,
		newYieldValuePositive,

		newTotalLabel,
		newTotalValue,
		newTotalValueNegative,
		newTotalValuePositive,
	} {
		style, err := doc.NewStyle(initer())
		if err != nil {
			return nil, err
		}
		styles = append(styles, style)
	}
	return &xlsxStyles{
		accName:         styles[0],
		colName:         styles[1],
		instrName:       styles[2],
		currencyGeneral: styles[3],
		general:         styles[4],
		borderFull:      styles[5],

		yieldLabel:         styles[6],
		yieldValue:         styles[7],
		yieldValueNegative: styles[8],
		yieldValuePositive: styles[9],

		totalLabel:         styles[10],
		totalValue:         styles[11],
		totalValueNegative: styles[12],
		totalValuePositive: styles[13],
	}, nil
}

func newAccNameStyle() *excelize.Style {
	return &excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 18.0},
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"ABD2BE"},
		},
		Border: newFullBorders(),
	}
}

func newColNamesStyle() *excelize.Style {
	return &excelize.Style{
		Font: &excelize.Font{Bold: true},
		Fill: excelize.Fill{
			Type:  "gradient",
			Color: []string{"F1F2CF", "ABD2BE"},
		},
		Border: newFullBorders(),
	}
}

func newInstrNameStyle() *excelize.Style {
	return &excelize.Style{
		Font: &excelize.Font{Bold: true},
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"F1CEFA"},
		},
		Border: newFullBorders(),
	}
}

// newCurrencyGeneralStyle Style with num format: #,##0.00
func newCurrencyGeneralStyle() *excelize.Style {
	return &excelize.Style{
		Border: newFullBorders(),
		NumFmt: 4, // #,##0.00
	}
}

func newGeneralStyle() *excelize.Style {
	return &excelize.Style{Border: newFullBorders()}
}

func newYieldLabel() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "gradient",
			Color: []string{"F1F2CF", "ABD2BE"},
		},
		Border: newFullBorders(),
	}
}

func newYieldValue() *excelize.Style {
	return &excelize.Style{
		Border: newFullBorders(),
		NumFmt: 2, // 0.00
	}
}

func newYieldValueNegative() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"FF9999"},
		},
		Border: newFullBorders(),
		NumFmt: 2,
	}
}

func newYieldValuePositive() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"90FA90"},
		},
		Border: newFullBorders(),
		NumFmt: 2,
	}
}

func newTotalLabel() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "gradient",
			Color: []string{"F1F2CF", "ABD2BE"},
		},
		Border: newFullBorders(),
	}
}

func newTotalValue() *excelize.Style {
	return &excelize.Style{
		Border: newFullBorders(),
		NumFmt: 2, // 0.00
	}
}

func newTotalValueNegative() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"FF9999"},
		},
		Border: newFullBorders(),
		NumFmt: 2,
	}
}

func newTotalValuePositive() *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"90FA90"},
		},
		Border: newFullBorders(),
		NumFmt: 2,
	}
}

func newFullBorders() []excelize.Border {
	return []excelize.Border{
		{Type: "left", Style: 2},
		{Type: "right", Style: 2},
		{Type: "top", Style: 2},
		{Type: "bottom", Style: 2},
	}
}

func newBorderFull() *excelize.Style {
	return &excelize.Style{Border: newFullBorders()}
}
