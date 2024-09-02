package entity

import "errors"

var (
	ErrTokenNotFound          = errors.New("токен не найден")
	ErrTokenExist             = errors.New("такой токен уже существует")
	ErrTokenIncorrectPassword = errors.New("неверный пароль шифрованного токена")
	ErrTokenValueNotSet       = errors.New("значение токена не было установлено")

	ErrInstrumentNotFound = errors.New("интструмент не найден")
)
