package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
)

var _ interfaces.Content = new(doc)

const docText = `
### Содержание

1. Цели программы
3. Основные положения
2. Руководство по использованию
`

type doc struct {
	content fyne.CanvasObject
}

func NewDoc(parentApp interfaces.App) *doc {
	return &doc{
		content: container.NewVBox(
			widget.NewRichTextFromMarkdown("## Документация"),
			widget.NewRichTextFromMarkdown(docText),
		),
	}
}

func (c *doc) Content() fyne.CanvasObject {
	return c.content
}

func (c *doc) Refresh() error {
	return nil
}
