package app

import (
	"context"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

	addToken            guiInterfaces.Content
	creation            guiInterfaces.Content
	about               guiInterfaces.Content
	progressBarInfinite guiInterfaces.ProgressBarInfinite

	aboutWindow fyne.Window
}

func NewApp(ctx context.Context, l logger.Logger, services guiInterfaces.AppServices) *application {
	a := app.NewWithID(gui.AppID)

	ctx, cancel := context.WithCancel(ctx)
	mw := newMainWindow(a, func() {
		l.Infof("Quit...")
		cancel()
		a.Quit()
	})

	guiApp := &application{
		fyneApp:    a,
		mainWindow: mw,
		logger:     l,
		ctx:        ctx,
		services:   services,
	}
	guiApp.setMainMenu()

	return guiApp
}

func newMainWindow(a fyne.App, onClosed func()) fyne.Window {
	w := a.NewWindow(gui.AppTitle)
	w.SetMaster()
	w.SetOnClosed(onClosed)
	w.Resize(gui.WinSize())
	w.SetFixedSize(gui.WinFixedSize)

	return w
}

func (a *application) ShowAndRun() error {
	var err error
	if err = a.Refresh(); err != nil {
		a.showFatal(err)
	}
	a.MainWindow().ShowAndRun()
	return err
}

func (a *application) Refresh() error {
	tokens, err := a.services.Token().ListTokens(a.ctx)
	if err != nil {
		a.Logger().Errorf("error on list tokens at sturtup: %s", err.Error())
		return err
	}

	show := a.ShowCreateReport
	if len(tokens) == 0 {
		show = a.ShowAddToken
	}

	if err := a.AddToken().Refresh(); err != nil {
		return err
	}
	if err := a.CreateReport().Refresh(); err != nil {
		return err
	}
	show()

	return nil
}

func (a *application) ShowAddToken() {
	a.AddToken().Refresh()
	a.MainWindow().SetContent(container.NewVBox(
		a.ProgressBarInfinite().Content(),
		a.AddToken().Content(),
	))
}

func (a *application) ShowCreateReport() {
	a.CreateReport().Refresh()
	a.MainWindow().SetContent(container.NewVBox(
		a.ProgressBarInfinite().Content(),
		a.CreateReport().Content(),
	))
}

func (a *application) ShowError(err error) {
	dialog.ShowError(err, a.mainWindow)
}

func (a *application) showFatal(err error) {
	w := a.MainWindow()
	w.SetTitle("Критическая Ошибка")
	w.SetIcon(theme.ErrorIcon())
	w.SetContent(container.NewBorder(
		widget.NewLabel(err.Error()),
		&widget.Button{
			Text:       "Выход",
			Icon:       theme.CancelIcon(),
			Importance: widget.DangerImportance,
			OnTapped:   w.Close,
		},
		nil,
		nil,
	))
	w.Show()
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

func (a *application) AddToken() guiInterfaces.Content {
	if a.addToken == nil {
		a.addToken = content.NewAddToken(a)
	}

	return a.addToken
}

func (a *application) CreateReport() guiInterfaces.Content {
	if a.creation == nil {
		a.creation = content.NewCreateReport(a)
	}

	return a.creation
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
	if a.progressBarInfinite == nil {
		a.progressBarInfinite = component.NewProgressBarInfinite()
	}

	return a.progressBarInfinite
}

func (a *application) setMainMenu() {
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Файл", &fyne.MenuItem{
			Label:  "Выход",
			IsQuit: true,
			Action: a.MainWindow().Close,
		}),
		fyne.NewMenu("Справка", &fyne.MenuItem{
			Label:  "О прорамме",
			Action: a.showAbout,
		}),
	)
	a.MainWindow().SetMainMenu(mainMenu)
}

func (a *application) showAbout() {
	if a.aboutWindow == nil {
		a.createAbout()
	}
	a.about.Refresh()
	a.aboutWindow.CenterOnScreen()
	a.aboutWindow.Show()
}

func (a *application) createAbout() {
	about := a.fyneApp.NewWindow("О программе")
	if a.about == nil {
		a.about = content.NewAbout(a)
	}
	about.SetContent(a.about.Content())
	about.SetFixedSize(true)
	about.SetOnClosed(func() {
		a.mainWindow.Show()
		about.Close()
		a.aboutWindow = nil
	})
	a.aboutWindow = about
}
