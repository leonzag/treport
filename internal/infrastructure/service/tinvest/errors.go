package tinvest

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrTokenInvalid        = errors.New("некорректный токен")
	ErrUnauthenticated     = errors.New("ошибка аутентификации")
	ErrAlreayAuthenticated = errors.New("уже есть активное соединение с другим токеном доступа")
	ErrPing                = errors.New("неудачная проверка доступности сервиса")
	ErrNotFound            = errors.New("не найдено")
)

func IsUnauthenticatedError(err error) bool {
	return status.Code(err) == codes.Unauthenticated
}

func IsTokenError(err error) bool {
	return IsUnauthenticatedError(err)
}

func NotFound(err error) bool {
	return status.Code(err) == codes.NotFound
}

func ParseError(err error) error {
	switch {
	case IsTokenError(err):
		return ErrUnauthenticated
	case IsUnauthenticatedError(err):
		return ErrUnauthenticated
	case NotFound(err):
		return ErrNotFound
	}
	return err
}
