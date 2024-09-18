package content

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/presentation/gui/component/button"
	"github.com/leonzag/treport/internal/presentation/gui/component/input"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

var _ interfaces.Content = new(addTokenContent)

type addTokenContent struct {
	title             *widget.Entry
	token             *input.EntryToken
	password          *input.PasswordChecked
	selectExistingBtn *widget.Button

	app     interfaces.App
	form    *widget.Form
	content fyne.CanvasObject
}

func NewAddToken(parentApp interfaces.App) *addTokenContent {
	c := &addTokenContent{app: parentApp}
	c.form = widget.NewForm()

	c.title = input.NewTitleEntry()
	c.token = input.NewTokenEntry()
	c.password = input.NewPasswordCheckedEntry()
	c.selectExistingBtn = button.NewSelectExistingToken(c.app.ShowCreateReport)

	c.title.Validator = validator.RequiredField
	c.token.Validator = validator.RequiredField
	c.password.Validator = validator.NewPasswordDefaultValidator()

	c.title.OnChanged = func(input string) {
		invalid := c.title.Validate() != nil
		tokenInvalid := c.token.Validate() != nil
		c.token.SetDisabled(invalid)
		c.password.SetDisabled(invalid || tokenInvalid)
		c.form.Refresh()
	}

	c.token.OnChanged = func(input string) {
		invalid := c.token.Validate() != nil
		c.password.SetDisabled(invalid)
		c.form.Refresh()
	}

	c.form.OnSubmit = c.addTokenClick
	c.form.SubmitText = "Создать"

	c.form.Append("Название", c.title)
	c.form.Append("Токен", c.token)
	c.form.Append("Шифрование", c.password)

	c.token.Disable()
	c.password.Disable()

	c.content = container.NewVBox(
		container.NewCenter(widget.NewRichTextFromMarkdown("# Добавить токен")),
		container.NewPadded(c.form),
		container.NewCenter(c.selectExistingBtn),
	)

	return c
}

func (c *addTokenContent) addTokenClick() {
	title, pwd := c.title.Text, c.password.Text
	if title == "" {
		c.app.ShowError(errors.Join(
			fmt.Errorf("Некоторые необходимые поля не заполнены."),
			c.form.Validate(),
		))
		return
	}
	tokenDTO := dto.TokenRequestDTO{
		Title:    title,
		Password: pwd,
		Token:    c.token.Text,
	}

	srv := c.app.Services().Token()
	ctx := c.app.Ctx()
	if _, err := srv.AddToken(ctx, tokenDTO); err != nil {
		c.app.ShowError(err)
		return
	}
	if err := c.app.Refresh(); err != nil {
		c.app.ShowError(err)
		return
	}
	c.app.ShowCreateReport()
}

func (c *addTokenContent) Content() fyne.CanvasObject {
	return c.content
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
	c.title.SetText("")
	c.title.Refresh()
	c.token.Clear()
	c.password.Clear()
}
