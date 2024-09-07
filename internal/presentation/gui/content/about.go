package content

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
)

var _ interfaces.Content = new(addTokenContent)

type about struct {
	tabs    *container.AppTabs
	content fyne.CanvasObject
}

const desc = `
## Описание

Генератор отчётов по инвестиционным портфелям в **T-Invest**.
`

const license = `
> Эта программа распространяется без каких-либо гарантий.
`

func NewAbout(parentApp interfaces.App) *about {
	copyToClipboard := func(s string) {
		parentApp.MainWindow().Clipboard().SetContent(s)
	}
	openURL := func(u *url.URL) {
		parentApp.OpenURL(u)
	}

	return &about{
		content: container.NewAppTabs(

			container.NewTabItem("О программе", container.NewVBox(
				widget.NewRichTextFromMarkdown(desc),
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
						copyToClipboard(gui.ContactMail)
						openURL(gui.ContactMailURL())
					},
				},

				widget.NewLabel("Репозиторий"),

				&widget.Button{
					Text:       "Открыть в браузере",
					Icon:       theme.SearchIcon(),
					Importance: widget.LowImportance,
					Alignment:  widget.ButtonAlignLeading,
					OnTapped: func() {
						copyToClipboard(gui.ContactRepoURL().String())
						openURL(gui.ContactRepoURL())
					},
				},
			)),
		),
	}
}

func (c *about) Content() fyne.CanvasObject {
	return c.content
}

func (c *about) Refresh() error {
	return nil
}
