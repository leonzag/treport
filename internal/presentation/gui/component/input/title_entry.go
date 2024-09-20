package input

import (
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

func NewTitleEntry() *widget.Entry {
	e := widget.NewEntry()
	e.Validator = validator.RequiredField
	return e
}
