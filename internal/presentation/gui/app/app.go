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
	"github.com/leonzag/treport/internal/presentation/gui/component/input"
	"github.com/leonzag/treport/internal/presentation/gui/content"
	guiInterfaces "github.com/leonzag/treport/internal/presentation/gui/interfaces"
	"github.com/leonzag/treport/pkg/logger"
)

// check interface impl.
var _ guiInterfaces.App = new(application)

type application struct {
	fyneApp     fyne.App
	mainWindow  fyne.Window
	childWindow fyne.Window
	mainMenu    *fyne.MainMenu
	logger      logger.Logger
	services    guiInterfaces.AppServices
	ctx         context.Context
	cancel      context.CancelFunc

	addToken            guiInterfaces.Content
	creation            guiInterfaces.Content
	about               guiInterfaces.Content
	doc                 guiInterfaces.Content
	progressBarInfinite guiInterfaces.ProgressBarInfinite
}

func NewApp(ctx context.Context, l logger.Logger, services guiInterfaces.AppServices) *application {
	a := app.NewWithID(gui.AppID)
	ctx, cancel := context.WithCancel(ctx)

	return &application{
		fyneApp:  a,
		logger:   l,
		services: services,
		ctx:      ctx,
		cancel:   cancel,
	}
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
		a.Logger().Errorf("error on tokens list on app refresh: %s", err.Error())
		return err
	}

	switch len(tokens) {
	case 0:
		a.ShowAddToken()
	default:
		a.ShowCreateReport()
	}

	return nil
}

func (a *application) setContent(c guiInterfaces.Content) {
	c.Refresh()
	content := container.NewVBox(
		a.ProgressBarInfinite().Content(),
		c.Content(),
	)
	a.MainWindow().SetContent(content)
}

func (a *application) ShowAddToken() {
	a.setContent(a.addTokenContent())
}

func (a *application) ShowCreateReport() {
	a.setContent(a.createReportContent())
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
	input.ShowPasswordDialog(a.mainWindow, title, onSubmit)
}

func (a *application) addTokenContent() guiInterfaces.Content {
	if a.addToken == nil {
		a.addToken = content.NewAddToken(a)
	}

	return a.addToken
}

func (a *application) createReportContent() guiInterfaces.Content {
	if a.creation == nil {
		a.creation = content.NewCreateReport(a)
	}

	return a.creation
}

func (a *application) MainWindow() fyne.Window {
	if a.mainWindow == nil {
		a.mainWindow = a.fyneApp.NewWindow(gui.AppTitle)
		a.mainWindow.SetMaster()
		a.mainWindow.SetOnClosed(a.Quit)
		a.mainWindow.SetMainMenu(a.MainMenu())
		a.mainWindow.SetFixedSize(gui.WinFixedSize)
		a.mainWindow.Resize(gui.WinSize())
	}

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

func (a *application) MainMenu() *fyne.MainMenu {
	fileItems := []*fyne.MenuItem{
		{Label: "Выход", IsQuit: true, Action: a.Quit},
	}
	helpItems := []*fyne.MenuItem{
		{Label: "Документация", Action: a.showDoc},
		{Label: "О прорамме", Action: a.showAbout},
	}

	return fyne.NewMainMenu(
		fyne.NewMenu("Файл", fileItems...),
		fyne.NewMenu("Справка", helpItems...),
	)
}

func (a *application) showWindow(title string, c guiInterfaces.Content) {
	if a.childWindow != nil {
		a.childWindow.Close()
	}
	a.childWindow = a.fyneApp.NewWindow(title)
	a.childWindow.SetFixedSize(true)
	a.childWindow.SetOnClosed(func() {
		a.MainWindow().Show()
		a.MainWindow().RequestFocus()
		a.childWindow.Close()
	})
	c.Refresh()
	a.childWindow.SetContent(c.Content())
	a.childWindow.Show()
}

func (a *application) showDoc() {
	if a.doc == nil {
		a.doc = content.NewDoc(a)
	}
	a.showWindow("Документация", a.doc)
}

func (a *application) showAbout() {
	if a.about == nil {
		a.about = content.NewAbout(a)
	}
	a.showWindow("О программе", a.about)
}

func (a *application) Quit() {
	a.Logger().Infof("Quit...")
	a.cancel()
	a.fyneApp.Quit()
}
