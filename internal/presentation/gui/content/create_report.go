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

var _ interfaces.Content = new(createReportContent)

type createReportContent struct {
	tokenSelect *widget.Select

	form *widget.Form
	app  interfaces.App
}

func NewCreateReport(parentApp interfaces.App) *createReportContent {
	c := &createReportContent{app: parentApp}

	c.tokenSelect = widget.NewSelect([]string{}, func(opt string) {})
	c.tokenSelect.SetSelectedIndex(0)
	tokenDelBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		c.showDeleteTokenDialog()
	})
	tokenDelBtn.Importance = widget.MediumImportance
	selectField := container.NewBorder(nil, nil, nil, tokenDelBtn, c.tokenSelect)

	toAddTokenBtn := widget.NewButtonWithIcon(
		"Добавить ноый токен",
		theme.ContentAddIcon(),
		func() { parentApp.ShowAddToken() },
	)
	toAddTokenBtn.Importance = widget.LowImportance
	toAddTokenBtn.Alignment = widget.ButtonAlignLeading

	c.form = &widget.Form{
		Items: []*widget.FormItem{
			{Widget: widget.NewRichTextFromMarkdown("# Создать отчет")},
			{Text: "Выбор Tinvest Токена", Widget: selectField},
			// {Widget: container.NewBorder(nil, nil, toAddTokenBtn, nil)},
			{Widget: container.NewStack(toAddTokenBtn)},
		},
		OnSubmit:   c.showCreatePortfolioDialog,
		SubmitText: "Создать отчет",
	}

	return c
}

func (c *createReportContent) Refresh() error {
	c.form.Refresh()

	ctx := c.app.Ctx()
	tokenSrv := c.app.Services().Token()
	titles, err := tokenSrv.ListTokensTitles(ctx)
	if err != nil {
		return err
	}
	c.tokenSelect.SetOptions(titles)
	c.tokenSelect.SetSelectedIndex(0)

	return nil
}

func (c *createReportContent) Content() fyne.CanvasObject {
	return c.form
}

func (c *createReportContent) showDeleteTokenDialog() {
	q, msg := "Удалить токен?", "Это действие необратимо."
	c.app.ShowConfirm(q, msg, func(confirm bool) {
		if !confirm {
			return
		}
		tokenSrv := c.app.Services().Token()
		token := dto.NewTokenDTO(c.tokenSelect.Selected, "", "", "")
		if err := tokenSrv.DeleteToken(c.app.Ctx(), token); err != nil {
			c.app.ShowError(err)
		}
		if err := c.app.Refresh(); err != nil {
			c.app.ShowError(err)
		}
	})
}

func (c *createReportContent) showCreatePortfolioDialog() {
	tokenSrv := c.app.Services().Token()

	title := c.tokenSelect.Selected
	token, err := tokenSrv.GetTokenByTitle(c.app.Ctx(), title)
	if err != nil {
		c.app.ShowError(err)
		return
	}

	if token.Password == "" {
		if !c.checkToken(token.Token) {
			return
		}
		c.selectFolderAndReport(token.Token)
		return
	}

	c.app.ShowPasswordEnter("Токен защищен паролем", func(pwd string) {
		token, err = tokenSrv.GetTokenByTitleDecrypted(c.app.Ctx(), title, pwd)
		if err != nil {
			if errors.Is(err, entity.ErrTokenIncorrectPassword) {
				c.app.Logger().Infof("%s", err.Error())
				c.app.ShowError(err)
				return
			}
			c.app.Logger().Errorf("%s", err.Error())
			c.app.ShowError(err)
			return
		}
		if !c.checkToken(token.Token) {
			return
		}
		c.selectFolderAndReport(token.Token)
	})
}

func (c *createReportContent) checkToken(token string) bool {
	c.app.ProgressBarInfinite().Show()
	defer c.app.ProgressBarInfinite().Hide()

	srv := c.app.Services().Portfolio()
	if err := srv.Ping(c.app.Ctx(), token); err != nil {
		c.app.Logger().Errorf("%s", err.Error())
		c.app.ShowError(err)
		return false
	}
	return true
}

func (c *createReportContent) selectFolderAndReport(token string) {
	c.app.ShowFolderOpen(func(dest string, err error) {
		c.app.ProgressBarInfinite().Show()
		defer c.app.ProgressBarInfinite().Hide()

		file, err := c.createReport(dest, token)
		if err != nil {
			c.app.Logger().Errorf("%s", err.Error())
			c.app.ShowError(err)
			return
		}
		msg := fmt.Sprintf("Файл отчета: %s", file)
		c.app.Logger().Infof("created report file: %s", file)
		c.app.ShowInfo("Готово", msg)
	})
}

func (c *createReportContent) createReport(dest string, token string) (string, error) {
	portfolioSrv := c.app.Services().Portfolio()
	portfolioSrv.SetUseCache(true)
	reportSrv := c.app.Services().Report()

	summaryCh := make(chan []*entity.PortfolioSummary)
	errCh := make(chan error)

	ctxWithTimeout, cancel := context.WithTimeout(c.app.Ctx(), 30*time.Second)
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

func (c *createReportContent) clear() {
	c.tokenSelect.SetSelectedIndex(0)
}
