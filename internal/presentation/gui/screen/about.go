package screen

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/config"
	"github.com/leonzag/treport/internal/presentation/gui"
)

var _ Screen = new(AboutScreen)

type AboutScreen struct {
	content fyne.CanvasObject
}

const desc = `
## Описание

Генератор отчётов по инвестиционным портфелям в **T-Invest**.
`

const license = `
> Эта программа распространяется без каких-либо гарантий.
`

func NewAbout(appHandlers AppHandlers) *AboutScreen {
	return &AboutScreen{
		content: container.NewAppTabs(

			container.NewTabItem("О программе", container.NewVBox(
				widget.NewRichTextFromMarkdown(desc),
				widget.NewRichTextFromMarkdown(fmt.Sprintf(
					"## Версия \n%s-**%s**",
					config.AppName,
					fyne.CurrentApp().Metadata().Version,
				)),
				widget.NewRichTextFromMarkdown(license),
			)),

			container.NewTabItem("Контакты", container.New(
				layout.NewFormLayout(),

				widget.NewLabel("Почта"),

				&widget.Button{
					Text:       gui.ContactMail,
					Icon:       theme.ContentCopyIcon(),
					Importance: widget.LowImportance,
					Alignment:  widget.ButtonAlignLeading,
					OnTapped: func() {
						appHandlers.ToClipboard(gui.ContactMail)
						fyne.CurrentApp().OpenURL(gui.ContactMailURL())
					},
				},

				widget.NewLabel("Репозиторий"),

				&widget.Button{
					Text:       "Открыть в браузере",
					Icon:       theme.SearchIcon(),
					Importance: widget.LowImportance,
					Alignment:  widget.ButtonAlignLeading,
					OnTapped: func() {
						appHandlers.ToClipboard(gui.ContactRepoURL().String())
						fyne.CurrentApp().OpenURL(gui.ContactRepoURL())
					},
				},
			)),
		),
	}
}

func (s *AboutScreen) Content() fyne.CanvasObject {
	return s.content
}

func (s *AboutScreen) Refresh() error {
	return nil
}
