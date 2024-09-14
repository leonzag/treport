package input

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui/validator"
)

func ShowPasswordDialog(w fyne.Window, title string, onSubmit func(pwd string)) {
	pwdBind := binding.NewString()

	entry := widget.NewPasswordEntry()
	entry.Bind(pwdBind)
	entry.SetPlaceHolder("Введите пароль")
	entry.Validator = validator.RequiredField

	form := []*widget.FormItem{{Widget: entry}}
	dialog.ShowForm(title, "Ok", "Abort", form, func(confirm bool) {
		if !confirm {
			return
		}
		pwd, err := pwdBind.Get()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		onSubmit(pwd)
	}, w)
}
