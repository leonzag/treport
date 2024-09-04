package content

import (
	"context"
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	"github.com/leonzag/treport/internal/presentation/gui/interfaces"
)

var _ interfaces.Content = new(createReportForm)

type createReportForm struct {
	tokenSelect *widget.Select

	form *widget.Form
	app  interfaces.App
}

func NewCreationPage(parentApp interfaces.App) *createReportForm {
	p := &createReportForm{app: parentApp}

	p.tokenSelect = widget.NewSelect([]string{}, func(opt string) {})
	p.tokenSelect.SetSelectedIndex(0)
	tokenDelBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		p.showDeleteTokenDialog()
	})
	tokenDelBtn.Importance = widget.MediumImportance
	selectField := container.NewBorder(nil, nil, nil, tokenDelBtn, p.tokenSelect)

	toAddTokenBtn := widget.NewButtonWithIcon(
		"Добавить ноый токен",
		theme.ContentAddIcon(),
		func() { parentApp.ShowAddToken() },
	)
	toAddTokenBtn.Importance = widget.LowImportance
	toAddTokenBtn.Alignment = widget.ButtonAlignLeading

	p.form = &widget.Form{
		Items: []*widget.FormItem{
			{Widget: widget.NewRichTextFromMarkdown("# Создать отчет")},
			{Text: "Выбор Tinvest Токена", Widget: selectField},
			// {Widget: container.NewBorder(nil, nil, toAddTokenBtn, nil)},
			{Widget: container.NewStack(toAddTokenBtn)},
		},
		OnSubmit:   p.showCreatePortfolioDialog,
		SubmitText: "Создать отчет",
	}

	return p
}

func (p *createReportForm) Refresh() error {
	p.form.Refresh()

	ctx := p.app.Ctx()
	tokenSrv := p.app.Services().Token()
	titles, err := tokenSrv.ListTokensTitles(ctx)
	if err != nil {
		return err
	}
	p.tokenSelect.SetOptions(titles)
	p.tokenSelect.SetSelectedIndex(0)

	return nil
}

func (p *createReportForm) Content() fyne.CanvasObject {
	return p.form
}

func (p *createReportForm) showDeleteTokenDialog() {
	q, msg := "Удалить токен?", "Это действие необратимо."
	p.app.ShowConfirm(q, msg, func(confirm bool) {
		if !confirm {
			return
		}
		tokenSrv := p.app.Services().Token()
		token := dto.NewTokenDTO(p.tokenSelect.Selected, "", "", "")
		if err := tokenSrv.DeleteToken(p.app.Ctx(), token); err != nil {
			p.app.ShowError(err)
		}
		if err := p.app.Refresh(); err != nil {
			p.app.ShowError(err)
		}
	})
}

func (p *createReportForm) showCreatePortfolioDialog() {
	tokenSrv := p.app.Services().Token()

	title := p.tokenSelect.Selected
	token, err := tokenSrv.GetTokenByTitle(p.app.Ctx(), title)
	if err != nil {
		p.app.ShowError(err)
		return
	}

	if token.Password == "" {
		if !p.checkToken(token.Token) {
			return
		}
		p.selectFolderAndReport(token.Token)
		return
	}

	p.app.ShowPasswordEnter("Токен защищен паролем", func(pwd string) {
		token, err = tokenSrv.GetTokenByTitleDecrypted(p.app.Ctx(), title, pwd)
		if err != nil {
			if errors.Is(err, entity.ErrTokenIncorrectPassword) {
				p.app.Logger().Infof("%s", err.Error())
				p.app.ShowError(err)
				return
			}
			p.app.Logger().Errorf("%s", err.Error())
			p.app.ShowError(err)
			return
		}
		if !p.checkToken(token.Token) {
			return
		}
		p.selectFolderAndReport(token.Token)
	})
}

func (p *createReportForm) checkToken(token string) bool {
	p.app.ProgressBarInfinite().Show()
	defer p.app.ProgressBarInfinite().Hide()

	srv := p.app.Services().Portfolio()
	if err := srv.Ping(p.app.Ctx(), token); err != nil {
		p.app.Logger().Errorf("%s", err.Error())
		p.app.ShowError(err)
		return false
	}
	return true
}

func (p *createReportForm) selectFolderAndReport(token string) {
	p.app.ShowFolderOpen(func(dest string, err error) {
		p.app.ProgressBarInfinite().Show()
		defer p.app.ProgressBarInfinite().Hide()

		file, err := p.createReport(dest, token)
		if err != nil {
			p.app.Logger().Errorf("%s", err.Error())
			p.app.ShowError(err)
			return
		}
		msg := fmt.Sprintf("Файл отчета: %s", file)
		p.app.Logger().Infof("created report file: %s", file)
		p.app.ShowInfo("Готово", msg)
	})
}

func (p *createReportForm) createReport(dest string, token string) (string, error) {
	portfolioSrv := p.app.Services().Portfolio()
	portfolioSrv.SetUseCache(true)
	reportSrv := p.app.Services().Report()

	summaryCh := make(chan []*entity.PortfolioSummary)
	errCh := make(chan error)

	ctxWithTimeout, cancel := context.WithTimeout(p.app.Ctx(), 30*time.Second)
	defer cancel()

	go func() {
		summary, err := portfolioSrv.SummaryAll(ctxWithTimeout, token)
		if err != nil {
			errCh <- err
			return
		}
		summaryCh <- summary
	}()

	select {
	case <-ctxWithTimeout.Done():
		return "", fmt.Errorf("Не удалось сделать отчет. Медленное соединение.")
	case err := <-errCh:
		return "", err
	case summary := <-summaryCh:
		for _, s := range summary {
			s.Portfolio.SortPositionsByTypes(
				enum.InstrumentType_SHARE,
				enum.InstrumentType_BOND,
				enum.InstrumentType_FUTURES,
				enum.InstrumentType_OPTION,
				enum.InstrumentType_ETF,
				enum.InstrumentType_INDEX,
				enum.InstrumentType_SP,
				enum.InstrumentType_CURRENCY,
				enum.InstrumentType_CLEARING_CERTIFICATE,
				enum.InstrumentType_COMMODITY,
				enum.InstrumentType_UNSPECIFIED,
			)
		}
		return reportSrv.CreateXLSX(dest, summary)
	}
}

func (p *createReportForm) clear() {
	p.tokenSelect.SetSelectedIndex(0)
}
