package content

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

var _ interfaces.Content = new(addTokenForm)

type addTokenForm struct {
	titleEntry         *widget.Entry
	tokenEntry         *widget.Entry
	encryptionCheck    *widget.Check
	encryptionPassword *widget.Entry
	selectExistingBtn  *widget.Button

	form *widget.Form
	app  interfaces.App
}

func NewAddTokenPage(parentApp interfaces.App) *addTokenForm {
	p := &addTokenForm{app: parentApp}

	p.titleEntry = widget.NewEntry()
	p.titleEntry.Validator = validator.RequiredField

	p.tokenEntry = widget.NewEntry()
	p.tokenEntry.Validator = validator.RequiredField
	p.tokenEntry.SetPlaceHolder("вставьте ваш InvestAPI токен")

	p.encryptionPassword = widget.NewPasswordEntry()
	p.encryptionPassword.SetPlaceHolder("введите пароль")
	p.encryptionPassword.Disable()
	p.encryptionCheck = widget.NewCheck("", func(checked bool) {
		switch {
		case checked:
			p.encryptionPassword.Enable()
		case !checked:
			p.encryptionPassword.SetText("")
			p.encryptionPassword.Disable()
		}
	})

	encryptionFields := container.NewBorder(nil, nil, p.encryptionCheck, nil, p.encryptionPassword)

	getFromTinvestBtn := widget.NewButtonWithIcon("Получить", theme.DownloadIcon(), func() {
		parentApp.OpenURL(&url.URL{
			Scheme: "https",
			Host:   "www.tbank.ru",
			Path:   "/invest/settings/api/",
		})
	})
	getFromTinvestBtn.Alignment = widget.ButtonAlignTrailing
	getFromTinvestBtn.Importance = widget.LowImportance

	tokenEntryFields := container.NewBorder(nil, nil, nil, getFromTinvestBtn, p.tokenEntry)

	p.selectExistingBtn = widget.NewButtonWithIcon(
		"Выбрать существующий",
		theme.SearchIcon(),
		p.selectExistingClick,
	)
	p.selectExistingBtn.Importance = widget.LowImportance

	p.form = &widget.Form{
		Items: []*widget.FormItem{
			{Widget: widget.NewRichTextFromMarkdown("# Добавить токен")},
			{Text: "Название", Widget: p.titleEntry},
			{Text: "Токен", Widget: tokenEntryFields},
			{Text: "Шифрование", Widget: encryptionFields},
			{Widget: container.NewBorder(p.selectExistingBtn, nil, nil, nil)},
		},
		SubmitText: "Создать",
		OnSubmit:   p.addTokenClick,
	}

	return p
}

func (p *addTokenForm) selectExistingClick() {
	p.app.ShowCreateReport()
}

func (p *addTokenForm) addTokenClick() {
	tokenDTO := dto.NewTokenDTO(
		p.titleEntry.Text,
		"",
		p.encryptionPassword.Text,
		p.tokenEntry.Text,
	)
	err := p.app.Services().Token().AddToken(p.app.Ctx(), tokenDTO)
	if err != nil {
		p.app.ShowError(err)
	} else if err := p.app.Refresh(); err != nil {
		p.app.ShowError(err)
	} else {
		p.app.ShowCreateReport()
	}
}

func (p *addTokenForm) Content() fyne.CanvasObject {
	return p.form
}

func (p *addTokenForm) Refresh() error {
	service, ctx := p.app.Services(), p.app.Ctx()
	tokens, err := service.Token().ListTokensTitles(ctx)
	if err == nil && len(tokens) > 0 {
		p.selectExistingBtn.Show()
	} else {
		p.selectExistingBtn.Hide()
	}
	p.clear()
	p.form.Refresh()

	return nil
}

func (p *addTokenForm) clear() {
	p.titleEntry.SetText("")
	p.titleEntry.Refresh()

	p.tokenEntry.SetText("")
	p.tokenEntry.Refresh()

	p.encryptionPassword.SetText("")
	p.encryptionPassword.Refresh()

	p.encryptionCheck.SetChecked(false)
	p.encryptionCheck.Refresh()
}
