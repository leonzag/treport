package app

import (
	"context"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/leonzag/treport/internal/presentation/gui"
	"github.com/leonzag/treport/internal/presentation/gui/component"
	"github.com/leonzag/treport/internal/presentation/gui/content"
	guiInterfaces "github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/pkg/logger"
)

// check interface impl.
var _ guiInterfaces.App = new(application)

type application struct {
	fyneApp    fyne.App
	mainWindow fyne.Window
	logger     logger.Logger
	services   guiInterfaces.AppServices
	ctx        context.Context

	addTokenPage        guiInterfaces.Content
	creationPage        guiInterfaces.Content
	progressBarInfinite guiInterfaces.ProgressBarInfinite
}

func NewApp(ctx context.Context, l logger.Logger, services guiInterfaces.AppServices) *application {
	fyneApp := app.NewWithID(gui.AppID)

	mainWindow := fyneApp.NewWindow(gui.AppTitle)
	mainWindow.Resize(fyne.Size{
		Width:  gui.WinWidth,
		Height: gui.WinHeight,
	})
	mainWindow.SetFixedSize(gui.WinFixedSize)

	guiApp := &application{
		fyneApp:    fyneApp,
		mainWindow: mainWindow,
		logger:     l,
		services:   services,
		ctx:        ctx,
	}
	guiApp.addTokenPage = content.NewAddTokenPage(guiApp)
	guiApp.creationPage = content.NewCreationPage(guiApp)

	guiApp.progressBarInfinite = component.NewProgressBarInfinite()

	return guiApp
}

func (a *application) ShowAndRun() error {
	var err error
	if err = a.Refresh(); err != nil {
		dlg := dialog.NewError(err, a.mainWindow)
		dlg.SetOnClosed(a.mainWindow.Close)
		dlg.Show()
	}
	a.mainWindow.ShowAndRun()
	return err
}

func (a *application) Refresh() error {
	showWindow := a.ShowCreateReport
	tokens, err := a.services.Token().ListTokens(a.ctx)
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		showWindow = a.ShowAddToken
	}
	if err := a.addTokenPage.Refresh(); err != nil {
		return err
	}
	if err := a.creationPage.Refresh(); err != nil {
		return err
	}
	showWindow()

	return nil
}

func (a *application) ShowAddToken() {
	a.addTokenPage.Refresh()
	content := container.NewVBox(
		a.progressBarInfinite.Content(),
		a.addTokenPage.Content(),
	)
	a.mainWindow.SetContent(content)
}

func (a *application) ShowCreateReport() {
	a.creationPage.Refresh()
	content := container.NewVBox(
		a.progressBarInfinite.Content(),
		a.creationPage.Content(),
	)
	a.mainWindow.SetContent(content)
}

func (a *application) ShowError(err error) {
	dialog.ShowError(err, a.mainWindow)
}

func (a *application) ShowInfo(title, msg string) {
	dialog.ShowInformation(title, msg, a.mainWindow)
}

func (a *application) ShowConfirm(title, msg string, callback func(bool)) {
	dialog.ShowConfirm(title, msg, callback, a.mainWindow)
}

func (a *application) ShowFolderOpen(callback func(string, error)) {
	dlg := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri == nil {
			return
		}
		callback(uri.Path(), err)
	}, a.mainWindow)
	dlg.SetConfirmText("Выбрать папку")
	dlg.SetDismissText("Закрыть")
	dlg.Show()
}

func (a *application) ShowPasswordEnter(title string, onSubmit func(pwd string)) {
	component.ShowPasswordDialog(a.mainWindow, title, onSubmit)
}

func (a *application) AddTokenForm() guiInterfaces.Content {
	return a.addTokenPage
}

func (a *application) CreateReportForm() guiInterfaces.Content {
	return a.creationPage
}

func (a *application) MainWindow() fyne.Window {
	return a.mainWindow
}

func (a *application) OpenURL(url *url.URL) {
	a.fyneApp.OpenURL(url)
}

func (a *application) Logger() logger.Logger {
	return a.logger
}

func (a *application) Services() guiInterfaces.AppServices {
	return a.services
}

func (a *application) Ctx() context.Context {
	return a.ctx
}

func (a *application) ProgressBarInfinite() guiInterfaces.ProgressBarInfinite {
	return a.progressBarInfinite
}
