package button

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui"
)

func NewGetTokenLinkButton() *widget.Button {
	return &widget.Button{
		Text:       "Получить",
		Icon:       theme.ContentAddIcon(),
		Importance: widget.LowImportance,
		Alignment:  widget.ButtonAlignTrailing,
		OnTapped: func() {
			fyne.CurrentApp().OpenURL(gui.GetTokenURL())
		},
	}
}
