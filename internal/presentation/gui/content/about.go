package content

import (
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

	aboutBox := container.NewVBox(
		widget.NewRichTextFromMarkdown(desc),
		widget.NewRichTextFromMarkdown(license),
	)

	mail := widget.NewButtonWithIcon(gui.ContactMail, theme.ContentCopyIcon(), func() {
		copyToClipboard(gui.ContactMail)
		parentApp.OpenURL(gui.ContactMailURL())
	})
	mail.Importance = widget.LowImportance
	mail.Alignment = widget.ButtonAlignLeading

	repo := widget.NewButtonWithIcon("Открыть в браузере", theme.SearchIcon(), func() {
		copyToClipboard(gui.ContactRepoURL().String())
		parentApp.OpenURL(gui.ContactRepoURL())
	})
	repo.Importance = widget.LowImportance
	repo.Alignment = widget.ButtonAlignLeading

	contacts := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Почта"),
		mail,
		widget.NewLabel("Репозиторий"),
		repo,
	)

	return &about{
		content: container.NewAppTabs(
			container.NewTabItem("О программе", aboutBox),
			container.NewTabItem("Контакты", contacts),
		),
	}
}

func (c *about) Content() fyne.CanvasObject {
	return c.content
}

func (c *about) Refresh() error {
	return nil
}
