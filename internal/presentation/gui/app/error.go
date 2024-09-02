package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/leonzag/treport/internal/presentation/gui"
)

func (a *application) showErrorDialogExit(err error, onClosed func()) {
	close := func() {
		onClosed()
		a.mainWindow.Close()
	}
	dlg := dialog.NewError(err, a.mainWindow)
	dlg.SetOnClosed(close)
	dlg.SetDismissText("Выход")
	dlg.Resize(fyne.NewSize(gui.WinWidth, gui.WinHeight))
	dlg.Refresh()
	dlg.Show()
}
