package input

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/component"
)

var (
	_ fyne.Widget      = new(PasswordChecked)
	_ fyne.Disableable = new(PasswordChecked)
)

type PasswordChecked struct {
	widget.BaseWidget
	entry *widget.Entry
	check *widget.Check

	// Do not use: required by widget.Form: for hint working
	Validator fyne.StringValidator `json:"-"`
	Text      string
	OnChanged func(string) `json:"-"`
}

func NewPasswordCheckedEntry() *PasswordChecked {
	w := &PasswordChecked{
		entry: widget.NewPasswordEntry(),
	}
	w.entry.OnChanged = func(input string) {
		w.Text = input
		if w.OnChanged != nil {
			w.OnChanged(input)
		}
	}
	w.check = widget.NewCheck("", w.onChecked)
	w.entry.Disable()
	w.ExtendBaseWidget(w)

	return w
}

func (w *PasswordChecked) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(component.NewHBoxTrailing(w.entry, w.check))
}

func (w *PasswordChecked) Validate() error {
	if w.Validator == nil {
		return nil
	}
	w.entry.Validator = w.Validator
	return w.entry.Validate()
}

func (w *PasswordChecked) SetOnValidationChanged(fn func(error)) {
	w.entry.SetOnValidationChanged(fn)
}

func (w *PasswordChecked) SetValidationError(err error) {
	w.entry.SetValidationError(err)
}

func (w *PasswordChecked) SetValidator(v fyne.StringValidator) {
	w.entry.Validator = v
	w.Validator = v
}

func (w *PasswordChecked) SetPlaceholder(text string) {
	w.entry.SetPlaceHolder(text)
}

func (w *PasswordChecked) onChecked(checked bool) {
	w.setInputEnabled(checked)
	w.entry.FocusGained()
}

func (w *PasswordChecked) setInputEnabled(enabled bool) {
	if enabled == !w.entry.Disabled() {
		return
	}
	action := map[bool]func(){
		false: w.entry.Disable,
		true:  w.entry.Enable,
	}[enabled]
	w.entry.SetText("")
	w.entry.SetValidationError(nil)
	action()
}

func (w *PasswordChecked) SetDisabled(disabled bool) {
	if w.Disabled() == disabled {
		return
	}

	action := map[bool]func(){
		false: w.Enable,
		true:  w.Disable,
	}[disabled]

	action()
}

func (w *PasswordChecked) Disabled() bool {
	return w.check.Disabled()
}

func (w *PasswordChecked) Disable() {
	w.Clear()
	w.check.Disable()
	w.entry.Disable()
	w.Refresh()
}

func (w *PasswordChecked) Enable() {
	w.check.Enable()
}

func (w *PasswordChecked) Clear() {
	w.Text = ""
	w.entry.Text = w.Text
	w.entry.SetValidationError(nil)
	w.check.SetChecked(false)
	w.Refresh()
}

func (w *PasswordChecked) Refresh() {
	w.entry.Text = w.Text
	w.entry.Refresh()
	w.check.Refresh()
}
