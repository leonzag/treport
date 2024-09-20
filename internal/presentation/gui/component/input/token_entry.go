package input

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/component"
	"github.com/leonzag/treport/internal/presentation/gui/component/button"
)

var (
	_ fyne.Widget      = new(EntryToken)
	_ fyne.Disableable = new(EntryToken)
)

type EntryToken struct {
	widget.BaseWidget
	entry       *widget.Entry
	getTokenBtn *widget.Button

	// Do not use: required by widget.Form: for hint working
	Validator fyne.StringValidator `json:"-"`
	Text      string
	OnChanged func(string) `json:"-"`
}

func NewTokenEntry() *EntryToken {
	w := &EntryToken{
		entry:       widget.NewEntry(),
		getTokenBtn: button.NewGetTokenLinkButton(),
	}
	w.entry.OnChanged = func(input string) {
		w.Text = input
		if w.OnChanged != nil {
			w.OnChanged(input)
		}
	}
	w.entry.PlaceHolder = "вставьте ваш InvestAPI токен"
	w.Validator = w.entry.Validator

	w.ExtendBaseWidget(w)

	return w
}

func (w *EntryToken) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(component.NewHBoxLeading(w.entry, w.getTokenBtn))
}

func (w *EntryToken) Validate() error {
	if w.Validator == nil {
		return nil
	}
	w.entry.Validator = w.Validator
	return w.entry.Validate()
}

func (w *EntryToken) SetOnValidationChanged(fn func(error)) {
	w.entry.SetOnValidationChanged(fn)
}

func (w *EntryToken) SetValidationError(err error) {
	w.entry.SetValidationError(err)
}

func (w *EntryToken) SetDisabled(disabled bool) {
	if w.Disabled() == disabled {
		return
	}
	action := map[bool]func(){
		false: w.Enable,
		true:  w.Disable,
	}[disabled]
	action()
}

func (w *EntryToken) Disabled() bool {
	return w.entry.Disabled()
}

func (w *EntryToken) Disable() {
	w.getTokenBtn.Disable()
	w.entry.Disable()
}

func (w *EntryToken) Enable() {
	w.getTokenBtn.Enable()
	w.entry.Enable()
}

func (w *EntryToken) Refresh() {
	w.entry.Refresh()
}

func (w *EntryToken) Clear() {
	w.entry.SetText("")
	w.SetValidationError(nil)
	w.Refresh()
}

func (w *EntryToken) Entry() *widget.Entry {
	return w.entry
}
