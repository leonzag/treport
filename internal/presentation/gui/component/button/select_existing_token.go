package button

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewSelectExistingToken(onTapped func()) *widget.Button {
	return &widget.Button{
		Text:       "Выбрать существующий",
		Icon:       theme.SearchIcon(),
		Importance: widget.LowImportance,
		OnTapped:   onTapped,
	}
}
