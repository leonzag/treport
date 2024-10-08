package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leonzag/treport/internal/presentation/gui"
)

var _ Screen = new(DocScreen)

var sections = []string{
	"# Документация",
	`## Содержание`,
	`
1. Цели программы
3. Основные положения
2. Руководство по использованию
`,
	`## Цели программы`,
	`
Иногда инвестицинные платформы
предоставляют возможность удобно
получить текущее состояние по текущим
портфелям.`,
	`
Это позволяет вести собственную аналитику.
`,
	`
Однако зачастую приходится кропотливо
переписывать показатели с экрана, а то и
искать необходимую информацию на разных
страницах.`,
	`
Поскольку нет удобного и быстрого способа
получить отчет по портфелям из мобильного
приложения или в браузере.`,
	`
Это как раз случай платформы T-Invest.`,
	`
Но T-Invest предоставляет удобное API для
разработчиков.`,
	`
Имея специальный ключ (токен) с помощью
запросов к API, можно получить всю
необходимую информацию.`,
	`
Цель этой программы в том, чтобы
автоматизировать процесс формирования
отчётов по состоянию портфелей.`,
	`
Она создаёт файл в формате Excel, в котором
на отдельном листе отражены ваши
финансовые инструменты и их состояние в портфеле, а также 
короткая промежуточная сводка: например
показатели доходности и другие.`,
	`## Основные положения`,
	`
Эту программу вы используете на свой страх
и риск.`,
	`
Программа взаимодействует с API T-Invest,
используя ваш токен доступа.`,
	`
Этот токен доступа вы не должны передавать
никому, поскольку он предназначен только
для личного использования.
`,
	`## Руководство по использованию`,
	`
Для того, чтобы использовать программу,
вам необходимо выпустить
**Токен Доступа API.**`,
	`
Сделать это можно по этой ссылке:
[https://www.tbank.ru/invest/settings/api/](tbank.ru/invest/settings/api)`,
	`
Этот токен вы добавляете в программу, и,
опционально, шифруете его на случай
несанкционированного доступа к вашему
компьютеру.`,
	`
Крайне желательно, чтобы ваш токен имел
доступ *"только для чтения"*, однако это
ваша личная отвественность, а не
требование программы.`,
	`
Далее, перейдя на экран **создания отчетов**,
вы можете создать документ с отчётом.`,
	`
Предварительно вам может потребоваться
ввести пароль, для того, чтобы
расшифровать токен доступа.`,
	`
По окончании процесса, файл Excel-документа
откроется в вашей программе для просмотра
документов по умолчанию.`,
	`
Если этого не произошло, но никаких ошибок
программа не выдала, то вы можете открыть
его вручную.`,
	`# Удачи!`,
}

type DocScreen struct {
	content fyne.CanvasObject
}

func NewDoc(appHandlers AppHandlers) *DocScreen {
	winSize := gui.WinSize().Max(fyne.NewSize(420, 240))
	body := container.NewVBox()
	for _, text := range sections {
		o := widget.NewRichTextFromMarkdown(text)
		o.Wrapping = fyne.TextWrapWord
		body.Add(o)
	}
	scroll := container.NewScroll(body)
	scroll.SetMinSize(winSize)
	return &DocScreen{
		content: scroll,
	}
}

func (s *DocScreen) Content() fyne.CanvasObject {
	return s.content
}

func (s *DocScreen) Refresh() error {
	return nil
}
