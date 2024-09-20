package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/presentation/gui/component"
	"github.com/leonzag/treport/internal/presentation/gui/component/button"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

var _ Screen = new(CreateReportScreen)

type CreateReportScreen struct {
	tokenSelect *widget.Select
	password    *widget.Entry

	app     AppHandlers
	form    *widget.Form
	content fyne.CanvasObject

	tokens map[string]*entity.Token
}

func NewCreateReport(appHandlers AppHandlers) *CreateReportScreen {
	s := &CreateReportScreen{
		app:  appHandlers,
		form: widget.NewForm(),
	}

	s.password = widget.NewPasswordEntry()
	s.password.SetPlaceHolder("введите пароль")
	s.password.Validator = validator.NewPasswordDefaultValidator()
	s.password.OnChanged = func(pwd string) {
		s.form.Validate()
	}

	s.tokenSelect = widget.NewSelect([]string{}, s.onSelectTokenClick)

	s.form.Append("T-Invest Токен", component.NewHBoxLeading(
		s.tokenSelect,
		button.NewDeleteToken(s.onDeleteTokenClick),
	))
	s.form.Append("Токен зашифрован", s.password)

	s.form.SubmitText = "Создать отчет"
	s.form.OnSubmit = s.onReportClick

	s.content = container.NewVBox(
		container.NewCenter(widget.NewRichTextFromMarkdown("# Создать отчет")),
		container.NewPadded(s.form),
		container.NewCenter(button.NewAddToken(s.onToScreenAddTokenClick)),
	)

	return s
}

func (s *CreateReportScreen) SetTokens(tokens map[string]*entity.Token) {
	s.tokens = tokens
	titles := make([]string, 0, len(tokens))
	for title := range s.tokens {
		titles = append(titles, title)
	}
	s.tokenSelect.SetOptions(titles)
	s.Refresh()
}

func (s *CreateReportScreen) Refresh() error {
	s.tokenSelect.SetSelectedIndex(0)
	s.form.Refresh()
	s.clear()

	return nil
}

func (s *CreateReportScreen) Content() fyne.CanvasObject {
	return s.content
}

func (s *CreateReportScreen) onSelectTokenClick(title string) {
	token, ok := s.tokens[title]
	if !ok {
		return
	}
	switch token.Password {
	case "":
		s.password.Disable()
		s.password.SetText("")
		s.password.SetValidationError(nil)
		s.password.Hide()
	default:
		s.password.Enable()
		s.password.Show()
		s.form.Validate()
	}
	s.form.Refresh()
}

func (s *CreateReportScreen) onReportClick() {
	s.app.CreateReport(dto.TokenRequestDTO{
		Title:    s.tokenSelect.Selected,
		Password: s.password.Text,
	})
	s.clear()
}

func (s *CreateReportScreen) onDeleteTokenClick() {
	s.app.DeleteToken(dto.TokenRequestDTO{
		Title: s.tokenSelect.Selected,
	})
}

func (s *CreateReportScreen) onToScreenAddTokenClick() {
	s.app.ToScreenAddToken()
}

func (s *CreateReportScreen) clear() {
	s.password.SetText("")
}
