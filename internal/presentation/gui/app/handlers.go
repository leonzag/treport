package app

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/skratchdot/open-golang/open"

	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/presentation/gui"
)

func (a *application) ToClipboard(s string) {
	a.mainWindow().Clipboard().SetContent(s)
}

func (a *application) ToScreenAddToken() {
	a.screenCreateReport().Refresh()
	a.showScreenAddToken()
}

func (a *application) ToScreenCreateReport() {
	a.screenAddToken().Refresh()
	a.showScreenCreateReport()
}

func (a *application) AddToken(tokenDTO dto.TokenRequestDTO) {
	if tokenDTO.Title == "" || tokenDTO.Token == "" {
		err := a.screenAddToken().Validate()
		err = fmt.Errorf("Необходимые поля не заполнены.\n%w", err)
		a.showError(err)
		return
	}

	var err error
	srv := a.Services().Token()
	a.withProgress(func() {
		_, err = srv.AddToken(a.Ctx(), tokenDTO)
	})
	if err != nil {
		a.showError(err)
		return
	}
	if err := a.Refresh(); err != nil {
		a.showError(err)
		return
	}
	a.screenAddToken().Refresh()
	a.showScreenCreateReport()
}

func (a *application) DeleteToken(token dto.TokenRequestDTO) {
	a.showConfirm("Удалить токен?", "Это действие необратимо.", func(confirm bool) {
		if !confirm {
			return
		}
		var err error
		a.withProgress(func() {
			err = a.Services().Token().DeleteToken(a.Ctx(), token)
		})
		if err != nil {
			a.Logger().Errorf("failed delete token '%s': %s", token.Title, err.Error())
			a.showError(err)
		}
		if err := a.Refresh(); err != nil {
			a.Logger().Errorf("failed refresh app after delete token: '%s'", err.Error())
			a.showError(err)
		}
	})
}

func (a *application) CreateReport(tokenDTO dto.TokenRequestDTO) {
	tokenSrv := a.Services().Token()
	var decrypted *entity.TokenDecrypted
	var err error

	token, ok := a.tokens[tokenDTO.Title]
	if !ok {
		err = fmt.Errorf("Не найден токен с названием: %s", tokenDTO.Title)
		a.Logger().Errorf(err.Error())
		a.showError(err)
		return
	}

	a.withProgress(func() {
		decrypted, err = tokenSrv.DecryptToken(token, tokenDTO.Password)
		if err != nil {
			return
		}
		if err = a.checkToken(decrypted.Token); err != nil {
			err = fmt.Errorf("Невалидный или просроченный токен!\n\n%w", err)
		}
	})
	if err != nil {
		a.Logger().Errorf(err.Error())
		a.showError(err)
		return
	}

	type folderOpen struct {
		err  error
		path string
	}

	folderCh := make(chan folderOpen)
	a.showFolderOpen(func(uri fyne.ListableURI, err error) {
		var path string
		if uri != nil {
			path = uri.Path()
		}
		folderCh <- folderOpen{err: err, path: path}
	})

	go func() {
		folder, ok := <-folderCh
		switch {
		case !ok:
			return
		case folder.err != nil:
			a.Logger().Errorf("%s", err.Error())
			a.showError(err)
			return
		case folder.path == "":
			return
		}
		a.Logger().Infof("selected folder:\n%s", folder.path)
		a.report(folder.path, decrypted.Token)
	}()
}

func (a *application) report(folder string, token string) {
	var err error
	var file string

	portfolioSrv := a.Services().Portfolio()
	portfolioSrv.SetUseCache(true)
	reportSrv := a.Services().Report()
	sorting := entity.NewPositionsSoringDefault()

	a.withProgress(func() {
		type summaryRes struct {
			err     error
			summary []*entity.PortfolioSummary
		}

		summaryResCh := make(chan summaryRes)
		ctx, cancel := context.WithTimeout(a.Ctx(), gui.TimeLimitLong)
		defer cancel()

		go func() {
			summary, err := portfolioSrv.SummaryAll(ctx, token)
			summaryResCh <- summaryRes{err: err, summary: summary}
		}()

		select {
		case <-ctx.Done():
			err = fmt.Errorf("Не удалось сделать отчет. Медленное соединение.")
		case res := <-summaryResCh:
			if res.err != nil {
				err = res.err
			}
			for _, s := range res.summary {
				s.Portfolio.SortPositionsByTypes(sorting...)
			}
			file, err = reportSrv.CreateXLSX(folder, res.summary)
		}
	})

	if err != nil {
		a.Logger().Errorf("%s", err.Error())
		a.showError(err)
		return
	}

	a.Logger().Infof("created report:\n%s", file)
	a.showInfo("Готово", fmt.Sprintf("Файл отчета: %s", file))

	if err := open.Start(file); err != nil {
		a.Logger().Errorf("failed to open file with default application: %s", file)
	}
}

func (a *application) checkToken(token string) error {
	srv := a.Services().Portfolio()
	return srv.Ping(a.Ctx(), token)
}
