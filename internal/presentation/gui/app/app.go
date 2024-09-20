package app

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/presentation/gui"
	"github.com/leonzag/treport/internal/presentation/gui/component"
	"github.com/leonzag/treport/internal/presentation/gui/component/input"
	"github.com/leonzag/treport/internal/presentation/gui/screen"
	"github.com/leonzag/treport/pkg/logger"
)

// check interface impl.
var _ App = new(application)

type application struct {
	fyneApp     fyne.App
	win         fyne.Window
	childWindow fyne.Window
	menu        *fyne.MainMenu
	logger      logger.Logger
	services    AppServices
	ctx         context.Context
	cancel      context.CancelFunc

	addToken *screen.AddTokenScreen
	creation *screen.CreateReportScreen
	about    *screen.AboutScreen
	doc      *screen.DocScreen
	progress *component.ProgressBarInfinite

	tokens          map[string]*entity.Token
	activeScreen    fyne.CanvasObject
	screenContainer *fyne.Container
}

func NewApp(ctx context.Context, l logger.Logger, services AppServices) *application {
	a := app.NewWithID(gui.AppID)
	ctx, cancel := context.WithCancel(ctx)

	return &application{
		fyneApp:  a,
		logger:   l,
		services: services,
		ctx:      ctx,
		cancel:   cancel,
		tokens:   make(map[string]*entity.Token),
	}
}

func (a *application) Ctx() context.Context {
	return a.ctx
}

func (a *application) Services() AppServices {
	return a.services
}

func (a *application) Logger() logger.Logger {
	return a.logger
}

func (a *application) ShowAndRun() error {
	var err error
	if err = a.Refresh(); err != nil {
		a.showFatal(err)
	}
	a.mainWindow().ShowAndRun()
	return err
}

func (a *application) Refresh() error {
	if err := a.loadTokens(); err != nil {
		return err
	}
	a.screenCreateReport().SetTokens(a.tokens)
	a.screenAddToken().SetTokens(a.tokens)

	switch len(a.tokens) {
	case 0:
		a.showScreenAddToken()
	default:
		a.showScreenCreateReport()
	}

	return nil
}

func (a *application) Quit() {
	a.Logger().Infof("Quit...")
	a.cancel()
	a.fyneApp.Quit()
}

func (a *application) mainWindow() fyne.Window {
	if a.win == nil {
		a.win = a.fyneApp.NewWindow(gui.AppTitle)
		a.win.SetMaster()
		a.win.SetOnClosed(a.Quit)
		a.win.SetMainMenu(a.mainMenu())
		a.win.SetFixedSize(gui.WinFixedSize)
		a.win.Resize(gui.WinSize())
	}

	return a.win
}

func (a *application) mainMenu() *fyne.MainMenu {
	if a.menu != nil {
		return a.menu
	}

	fileItems := []*fyne.MenuItem{
		{Label: "Выход", IsQuit: true, Action: a.Quit},
	}
	helpItems := []*fyne.MenuItem{
		{Label: "Документация", Action: a.showDoc},
		{Label: "О прорамме", Action: a.showAbout},
	}
	a.menu = fyne.NewMainMenu(
		fyne.NewMenu("Файл", fileItems...),
		fyne.NewMenu("Справка", helpItems...),
	)

	return a.menu
}

func (a *application) screenAddToken() *screen.AddTokenScreen {
	if a.addToken == nil {
		a.addToken = screen.NewAddToken(a)
	}

	return a.addToken
}

func (a *application) screenCreateReport() *screen.CreateReportScreen {
	if a.creation == nil {
		a.creation = screen.NewCreateReport(a)
	}

	return a.creation
}

func (a *application) progressBar() *component.ProgressBarInfinite {
	if a.progress == nil {
		a.progress = component.NewProgressBarInfinite()
	}

	return a.progress
}

func (a *application) loadTokens() error {
	tokens, err := a.services.Token().ListTokens(a.ctx)
	if err != nil {
		a.Logger().Errorf("error on tokens list on app refresh: %s", err.Error())
		return err
	}

	a.tokens = make(map[string]*entity.Token, len(tokens))
	for _, token := range tokens {
		a.tokens[token.Title] = token
	}

	return nil
}

func (a *application) setScreen(s screen.Screen) {
	if a.screenContainer == nil {
		a.screenContainer = container.NewPadded()
	}
	if a.activeScreen == nil {
		a.activeScreen = container.NewVBox(
			a.progressBar(),
			a.screenContainer,
		)
	}
	a.screenContainer.Objects = []fyne.CanvasObject{s.Content()}
	a.activeScreen.Refresh()
	a.mainWindow().SetContent(a.activeScreen)
}

func (a *application) showWindow(title string, s screen.Screen) {
	if a.childWindow != nil {
		a.childWindow.Close()
	}
	a.childWindow = a.fyneApp.NewWindow(title)
	a.childWindow.SetFixedSize(true)
	a.childWindow.SetOnClosed(func() {
		a.mainWindow().Show()
		a.mainWindow().RequestFocus()
		a.childWindow.Close()
	})
	s.Refresh()
	a.childWindow.SetContent(s.Content())
	a.childWindow.Show()
}

func (a *application) showDoc() {
	if a.doc == nil {
		a.doc = screen.NewDoc(a)
	}
	a.showWindow("Документация", a.doc)
}

func (a *application) showAbout() {
	if a.about == nil {
		a.about = screen.NewAbout(a)
	}
	a.showWindow("О программе", a.about)
}

func (a *application) showScreenAddToken() {
	a.setScreen(a.screenAddToken())
}

func (a *application) showScreenCreateReport() {
	a.setScreen(a.screenCreateReport())
}

func (a *application) showInfo(title, msg string) {
	dialog.ShowInformation(title, msg, a.win)
}

func (a *application) showConfirm(title, msg string, callback func(bool)) {
	dlg := dialog.NewConfirm(title, msg, callback, a.win)
	dlg.SetConfirmText("Подтвердить")
	dlg.SetDismissText("Отмена")
	dlg.Show()
}

func (a *application) showFolderOpen(callback func(string, error)) {
	dlg := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri == nil {
			return
		}
		callback(uri.Path(), err)
	}, a.win)
	dlg.SetConfirmText("Выбрать папку")
	dlg.SetDismissText("Закрыть")
	dlg.Show()
}

func (a *application) showPasswordEnter(title string, onSubmit func(pwd string)) {
	input.ShowPasswordDialog(a.win, title, onSubmit)
}

func (a *application) showError(err error) {
	dialog.ShowError(err, a.win)
}

func (a *application) showFatal(err error) {
	w := a.mainWindow()
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

func (a *application) withProgress(process func()) {
	a.progressBar().Show()
	defer a.progressBar().Hide()

	process()
}
