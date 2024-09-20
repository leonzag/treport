package app

import (
	"fmt"

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
	a.selectFolderAndReport(decrypted.Token)
}

func (a *application) selectFolderAndReport(token string) {
	a.showFolderOpen(func(dest string, err error) {
		var file string
		a.withProgress(func() {
			file, err = a.reportProcess(dest, token, gui.TimeLimitLong)
		})
		if err != nil {
			a.Logger().Errorf("%s", err.Error())
			a.showError(err)
			return
		}
		a.Logger().Infof("created report file: %s", file)
		a.showInfo("Готово", fmt.Sprintf("Файл отчета: %s", file))
		if err := open.Start(file); err != nil {
			a.Logger().Errorf("failed to open file with default application: %s", file)
		}
	})
}

func (a *application) checkToken(token string) error {
	srv := a.Services().Portfolio()
	return srv.Ping(a.Ctx(), token)
}
