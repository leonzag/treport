package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ Screen = new(DocScreen)

const docText = `
### Содержание

1. Цели программы
3. Основные положения
2. Руководство по использованию
`

type DocScreen struct {
	content fyne.CanvasObject
}

func NewDoc(appHandlers AppHandlers) *DocScreen {
	return &DocScreen{
		content: container.NewVBox(
			widget.NewRichTextFromMarkdown("## Документация"),
			widget.NewRichTextFromMarkdown(docText),
		),
	}
}

func (s *DocScreen) Content() fyne.CanvasObject {
	return s.content
}

func (s *DocScreen) Refresh() error {
	return nil
}
