package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
)

const (
	AppID       = "com.leonzag.treport"
	AppTitle    = "T-Invest: Портфельный Отчет"
	ContactMail = "leonzag997@gmail.com"

	WinWidth     = 720
	WinHeight    = 480
	WinFixedSize = true
)

func WinSize() fyne.Size {
	return fyne.Size{
		Width:  WinWidth,
		Height: WinHeight,
	}
}

func ContactRepoURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "www.github.com",
		Path:   "leonzag/treport",
	}
}

func ContactMailURL() *url.URL {
	return &url.URL{
		Scheme: "mailto",
		Host:   "leonzag997@gmail.com",
	}
}

func GetTokenURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "www.tbank.ru",
		Path:   "/invest/settings/api/",
	}
}
