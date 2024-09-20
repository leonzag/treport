package button

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewAddToken(onTapped func()) *widget.Button {
	return &widget.Button{
		Text:       "Добавить ноый токен",
		Icon:       theme.ContentAddIcon(),
		Importance: widget.LowImportance,
		Alignment:  widget.ButtonAlignLeading,
		OnTapped:   onTapped,
	}
}
