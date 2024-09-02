package report

import "github.com/xuri/excelize/v2"

type cellCtx struct {
	doc      *excelize.File
	sheet    string
	name     string
	col, row int
}

func (c *cellCtx) String() string {
	return c.name
}
