package dto

type TokenDTO struct {
	Title       string
	Description string
	Password    string
	Token       string
}

type TokenRequestDTO struct {
	Title       string
	Description string
	Password    string
	Token       string
}

func NewTokenDTO(title, desc, pwd, token string) TokenDTO {
	return TokenDTO{
		Title:       title,
		Description: desc,
		Password:    pwd,
		Token:       token,
	}
}
