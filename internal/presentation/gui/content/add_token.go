package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/presentation/gui"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

var _ interfaces.Content = new(addTokenContent)

type addTokenContent struct {
	titleEntry         *widget.Entry
	tokenEntry         *widget.Entry
	encryptionCheck    *widget.Check
	encryptionPassword *widget.Entry
	selectExistingBtn  *widget.Button

	form *widget.Form
	app  interfaces.App
}

func NewAddToken(parentApp interfaces.App) *addTokenContent {
	c := &addTokenContent{app: parentApp}

	c.titleEntry = widget.NewEntry()
	c.titleEntry.Validator = validator.RequiredField

	c.tokenEntry = widget.NewEntry()
	c.tokenEntry.Validator = validator.RequiredField
	c.tokenEntry.SetPlaceHolder("вставьте ваш InvestAPI токен")

	c.encryptionPassword = widget.NewPasswordEntry()
	c.encryptionPassword.SetPlaceHolder("введите пароль")
	c.encryptionPassword.Disable()
	c.encryptionCheck = widget.NewCheck("", func(checked bool) {
		switch {
		case checked:
			c.encryptionPassword.Enable()
		case !checked:
			c.encryptionPassword.SetText("")
			c.encryptionPassword.Disable()
		}
	})

	encryptionFields := container.NewBorder(nil, nil, c.encryptionCheck, nil, c.encryptionPassword)

	getNewTokenBtn := &widget.Button{
		Text:       "Получить",
		Icon:       theme.ContentAddIcon(),
		Importance: widget.LowImportance,
		Alignment:  widget.ButtonAlignTrailing,
		OnTapped: func() {
			parentApp.OpenURL(gui.GetTokenURL())
		},
	}

	tokenEntryFields := container.NewBorder(nil, nil, nil, getNewTokenBtn, c.tokenEntry)

	c.selectExistingBtn = &widget.Button{
		Text:       "Выбрать существующий",
		Icon:       theme.SearchIcon(),
		Importance: widget.LowImportance,
		OnTapped:   c.selectExistingClick,
	}

	c.form = &widget.Form{
		Items: []*widget.FormItem{
			{Widget: widget.NewRichTextFromMarkdown("# Добавить токен")},
			{Text: "Название", Widget: c.titleEntry},
			{Text: "Токен", Widget: tokenEntryFields},
			{Text: "Шифрование", Widget: encryptionFields},
			{Widget: container.NewBorder(c.selectExistingBtn, nil, nil, nil)},
		},
		SubmitText: "Создать",
		OnSubmit:   c.addTokenClick,
	}

	return c
}

func (c *addTokenContent) selectExistingClick() {
	c.app.ShowCreateReport()
}

func (c *addTokenContent) addTokenClick() {
	tokenDTO := dto.NewTokenDTO(
		c.titleEntry.Text,
		"",
		c.encryptionPassword.Text,
		c.tokenEntry.Text,
	)
	err := c.app.Services().Token().AddToken(c.app.Ctx(), tokenDTO)
	if err != nil {
		c.app.ShowError(err)
	} else if err := c.app.Refresh(); err != nil {
		c.app.ShowError(err)
	} else {
		c.app.ShowCreateReport()
	}
}

func (c *addTokenContent) Content() fyne.CanvasObject {
	return c.form
}

func (c *addTokenContent) Refresh() error {
	service, ctx := c.app.Services(), c.app.Ctx()
	tokens, err := service.Token().ListTokensTitles(ctx)
	if err == nil && len(tokens) > 0 {
		c.selectExistingBtn.Show()
	} else {
		c.selectExistingBtn.Hide()
	}
	c.clear()
	c.form.Refresh()

	return nil
}

func (c *addTokenContent) clear() {
	c.titleEntry.SetText("")
	c.titleEntry.Refresh()

	c.tokenEntry.SetText("")
	c.tokenEntry.Refresh()

	c.encryptionPassword.SetText("")
	c.encryptionPassword.Refresh()

	c.encryptionCheck.SetChecked(false)
	c.encryptionCheck.Refresh()
}
