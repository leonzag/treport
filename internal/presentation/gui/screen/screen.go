package screen

import (
	"fyne.io/fyne/v2"
	"github.com/leonzag/treport/internal/application/dto"
)

type Screen interface {
	Content() fyne.CanvasObject
	Refresh() error
}

type AppHandlers interface {
	ToScreenCreateReport()
	ToScreenAddToken()
	ToClipboard(text string)
	AddToken(token dto.TokenRequestDTO)
	CreateReport(token dto.TokenRequestDTO)
	DeleteToken(token dto.TokenRequestDTO)
}
