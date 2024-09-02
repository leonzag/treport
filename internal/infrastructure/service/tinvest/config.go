package tinvest

import "github.com/russianinvestments/invest-api-go-sdk/investgo"

func NewConfig(appName string, token string) investgo.Config {
	return investgo.Config{
		EndPoint: "invest-public-api.tinkoff.ru:443",
		AppName:  appName,
		Token:    token,
	}
}

func NewConfigSandbox(appName string, token string) investgo.Config {
	return investgo.Config{
		AppName: appName + "_test",
		Token:   token,
	}
}
