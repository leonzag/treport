package screen

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/presentation/gui/component/button"
	"github.com/leonzag/treport/internal/presentation/gui/component/input"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

var _ Screen = new(AddTokenScreen)

type appBind interface{}

type AddTokenScreen struct {
	title             *widget.Entry
	token             *input.EntryToken
	password          *input.PasswordChecked
	selectExistingBtn *widget.Button

	app     AppHandlers
	form    *widget.Form
	content fyne.CanvasObject

	tokens map[string]*entity.Token
}

func NewAddToken(appHandlers AppHandlers) *AddTokenScreen {
	s := &AddTokenScreen{
		app:    appHandlers,
		form:   widget.NewForm(),
		tokens: make(map[string]*entity.Token),

		title:    input.NewTitleEntry(),
		token:    input.NewTokenEntry(),
		password: input.NewPasswordCheckedEntry(),
	}
	s.selectExistingBtn = button.NewSelectExistingToken(s.onSelectExistingClick)

	s.title.Validator = validator.RequiredField
	s.token.Validator = validator.RequiredField
	s.password.Validator = validator.NewPasswordDefaultValidator()

	s.title.OnChanged = func(input string) {
		invalid := s.title.Validate() != nil
		tokenInvalid := s.token.Validate() != nil
		s.token.SetDisabled(invalid)
		s.password.SetDisabled(invalid || tokenInvalid)
		if _, ok := s.tokens[input]; ok {
			s.title.SetValidationError(fmt.Errorf("уже существует с таким именем"))
		}
		s.form.Refresh()
	}

	s.token.OnChanged = func(input string) {
		invalid := s.token.Validate() != nil
		s.password.SetDisabled(invalid)
		s.form.Refresh()
	}

	s.form.OnSubmit = s.onAddTokenClick
	s.form.SubmitText = "Создать"

	s.form.Append("Название", s.title)
	s.form.Append("Токен", s.token)
	s.form.Append("Шифрование", s.password)

	s.token.Disable()
	s.password.Disable()

	s.content = container.NewVBox(
		container.NewCenter(widget.NewRichTextFromMarkdown("# Добавить токен")),
		container.NewPadded(s.form),
		container.NewCenter(s.selectExistingBtn),
	)

	return s
}

func (s *AddTokenScreen) SetTokens(tokens map[string]*entity.Token) {
	s.tokens = tokens
	switch len(tokens) {
	case 0:
		s.selectExistingBtn.Disable()
	default:
		s.selectExistingBtn.Enable()
	}
}

func (s *AddTokenScreen) onAddTokenClick() {
	s.app.AddToken(dto.TokenRequestDTO{
		Title:    s.title.Text,
		Password: s.password.Text,
		Token:    s.token.Text,
	})
}

func (s *AddTokenScreen) onSelectExistingClick() {
	s.app.ToScreenCreateReport()
}

func (s *AddTokenScreen) Validate() error {
	return s.form.Validate()
}

func (s *AddTokenScreen) Content() fyne.CanvasObject {
	return s.content
}

func (s *AddTokenScreen) Refresh() error {
	s.clear()
	s.form.Refresh()

	return nil
}

func (s *AddTokenScreen) clear() {
	s.title.SetText("")
	s.title.SetValidationError(nil)
	s.title.Refresh()
	s.token.Clear()
	s.token.Refresh()
	s.password.Clear()
}
