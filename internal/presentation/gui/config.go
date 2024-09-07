package gui

import "net/url"

const (
	AppID       = "com.leonzag.treport"
	AppTitle    = "T-Invest: Портфельный Отчет"
	ContactMail = "leonzag997@gmail.com"

	WinWidth     = 720
	WinHeight    = 480
	WinFixedSize = true
)

var ContactRepoURL = url.URL{
	Scheme: "https",
	Host:   "www.github.com",
	Path:   "leonzag/treport",
}

var GetTokenURL = url.URL{
	Scheme: "https",
	Host:   "www.tbank.ru",
	Path:   "/invest/settings/api/",
}
