package report

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/pkg/logger"
	"github.com/xuri/excelize/v2"
)

type PortfolioReportCtl struct {
	logger logger.Logger
	file   *excelize.File
	sheet  string
	styles *xlsxStyles
}

func NewPortfolioRerpotController(l logger.Logger) *PortfolioReportCtl {
	return &PortfolioReportCtl{logger: l}
}

func (c *PortfolioReportCtl) Report(dest string, portfolios []*entity.PortfolioSummary) (string, error) {
	if err := checkDestFolder(dest); err != nil {
		c.logger.Errorf("report portfolio failed: %s", err.Error())
		return "", err
	}

	c.file = excelize.NewFile()
	defer c.file.Close()

	c.sheet = c.file.GetSheetName(0)
	if err := c.initStyles(); err != nil {
		c.logger.Errorf("report portfolio failed: %s", err.Error())
		return "", err
	}

	if err := c.newSummaryAll(portfolios); err != nil {
		c.logger.Errorf("report portfolio failed: %s", err.Error())
		return "", err
	}

	c.file.DeleteSheet("Sheet1")
	c.file.SetActiveSheet(0)

	return c.save(dest)
}

func (c *PortfolioReportCtl) newSummaryAll(portfolios []*entity.PortfolioSummary) error {
	for _, p := range portfolios {
		if err := c.newSummarySheet(p); err != nil {
			return err
		}
	}
	return nil
}

func (c *PortfolioReportCtl) newSummarySheet(p *entity.PortfolioSummary) error {
	c.sheet = p.Account.Name
	titles := summaryTitles()
	leftCell := "A1"
	rightCell, _ := excelize.CoordinatesToCellName(len(titles), 1)

	if _, err := c.file.NewSheet(c.sheet); err != nil {
		return err
	}

	if err := c.file.SetCellStr(c.sheet, leftCell, c.sheet); err != nil {
		return err
	}
	c.file.MergeCell(c.sheet, leftCell, rightCell)
	c.file.SetCellStyle(c.sheet, leftCell, rightCell, c.styles.accName)

	leftCell = "A2"
	rightCell, _ = excelize.CoordinatesToCellName(len(titles), 2)

	for i, title := range titles {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		if err := c.file.SetCellStr(c.sheet, cell, title); err != nil {
			return err
		}
	}
	c.file.SetCellStyle(c.sheet, leftCell, rightCell, c.styles.colName)

	row := 3
	cells := make([]string, len(p.Portfolio.Positions))
	for i := range cells {
		cells[i], _ = excelize.CoordinatesToCellName(1, row+i)
	}

	for i, cell := range cells {
		pos := p.Portfolio.Positions[i]
		instr, err := p.InstrumentByUid(pos.InstrumentUid)
		if err != nil {
			c.logger.Warnf("instrument info not found: %s", err.Error())
			instr = nil
		}

		if err := c.appendInstrumentPosition(cell, instr, pos); err != nil {
			return err
		}
	}
	row = row + len(cells) + 1

	yieldLabelCell, _ := excelize.JoinCellName("A", row)
	yieldValueCell, _ := excelize.JoinCellName("B", row)
	totalLabelCell, _ := excelize.JoinCellName("A", row+1)
	totalValueCell, _ := excelize.JoinCellName("B", row+1)

	c.file.SetCellStr(c.sheet, yieldLabelCell, "Текущая доходность портфеля, %:")
	c.file.SetCellFloat(c.sheet, yieldValueCell, p.Portfolio.ExpectedYield.ToFloat(), 2, 64)
	c.file.SetCellStyle(c.sheet, yieldLabelCell, yieldLabelCell, c.styles.yieldLabel)
	c.file.SetCellStyle(c.sheet, yieldValueCell, yieldValueCell, c.styles.yieldValue)

	c.file.SetCellStr(c.sheet, totalLabelCell, "Общая стоимость портфеля:")
	c.file.SetCellFloat(c.sheet, totalValueCell, p.Portfolio.TotalAmountPortfolio.ToFloat(), 2, 64)
	c.file.SetCellStyle(c.sheet, totalLabelCell, totalLabelCell, c.styles.totalLabel)
	c.file.SetCellStyle(c.sheet, totalValueCell, totalValueCell, c.styles.totalValue)

	c.fitCols()

	return nil
}

func (s *PortfolioReportCtl) appendInstrumentPosition(
	cell string,
	instr *entity.Instrument,
	pos *entity.PortfolioPosition,
) error {
	name := pos.InstrumentUid
	ticker := ""
	instrType := pos.InstrumentType
	if instr != nil {
		name, ticker, instrType = instr.Name, instr.Ticker, instr.InstrumentType
	}
	content := []interface{}{
		name,
		ticker,
		instrType,
		pos.QuantityLots.ToFloat(),
		pos.Quantity.ToFloat(),
		pos.AveragePositionPrice.ToFloat(),
		pos.AveragePositionPriceFifo.ToFloat(),
		pos.CurrentPrice.ToFloat(),
		pos.ExpectedYield.ToFloat(),
		pos.CurrentNkd.ToFloat(),
		pos.AveragePositionPricePt.ToFloat(),
		pos.ExpectedYieldFifo.ToFloat(),
		pos.Blocked,
		pos.BlockedLots.ToFloat(),
		pos.VarMargin.ToFloat(),
	}
	if err := s.file.SetSheetRow(s.sheet, cell, &content); err != nil {
		return err
	}

	col, row, _ := excelize.CellNameToCoordinates(cell)
	rightCell, _ := excelize.CoordinatesToCellName(col+14, row)
	s.file.SetCellStyle(s.sheet, cell, rightCell, s.styles.general)

	crcCellLeft, _ := excelize.CoordinatesToCellName(col+5, row)
	crcCellRight, _ := excelize.CoordinatesToCellName(col+10, row)
	s.file.SetCellStyle(s.sheet, cell, cell, s.styles.instrName)
	s.file.SetCellStyle(s.sheet, crcCellLeft, crcCellRight, s.styles.currencyGeneral)

	return nil
}

func (s *PortfolioReportCtl) initStyles() error {
	var err error
	s.styles, err = newStylesAll(s.file)

	return err
}

// fitCols Autofit all columns according to their text content.
func (s *PortfolioReportCtl) fitCols() {
	cols, err := s.file.GetCols(s.sheet)
	if err != nil {
		s.logger.Errorf("report portfolio: fit cols failed: %s", err.Error())
		return
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			s.logger.Errorf("report portfolio: fit cols failed: %s", err.Error())
			return
		}
		s.file.SetColWidth(s.sheet, name, name, float64(largestWidth))
	}
}

// save Save xlsx file and returns path to it.
func (s *PortfolioReportCtl) save(dest string) (string, error) {
	timestamp := time.Now().Format(time.RFC3339)
	if runtime.GOOS == "windows" {
		timestamp = strings.ReplaceAll(timestamp, ":", "-")
	}
	name := filepath.Join(dest, fmt.Sprintf("report-%s.xlsx", timestamp))
	if err := s.file.SaveAs(name); err != nil {
		s.logger.Errorf("report portfolio: save file failed: %s", err.Error())
		return "", err
	}

	return name, nil
}
