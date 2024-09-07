package interfaces

import (
	"context"
	"net/url"

	"fyne.io/fyne/v2"
	"github.com/leonzag/treport/pkg/logger"
)

type AppServices interface {
	Token() TokenService
	Portfolio() PortfolioService
	Report() PortfolioReportService
}

type App interface {
	Ctx() context.Context
	Services() AppServices
	Logger() logger.Logger

	MainWindow() fyne.Window
	ProgressBarInfinite() ProgressBarInfinite
	Refresh() error

	AddToken() Content
	ShowAddToken()

	CreateReport() Content
	ShowCreateReport()

	OpenURL(*url.URL)
	ShowInfo(title, msg string)
	ShowError(err error)
	ShowConfirm(title, msg string, callback func(confirm bool))
	ShowFolderOpen(callback func(folder string, err error))
	ShowPasswordEnter(title string, onSubmit func(pwd string))
}

type Content interface {
	Content() fyne.CanvasObject
	Refresh() error
}
