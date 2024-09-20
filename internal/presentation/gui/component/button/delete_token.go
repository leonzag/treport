package button

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewDeleteToken(onTapped func()) *widget.Button {
	return &widget.Button{
		Icon:       theme.DeleteIcon(),
		Importance: widget.MediumImportance,
		OnTapped:   onTapped,
	}
}
