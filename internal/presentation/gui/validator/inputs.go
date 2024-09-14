package validator

import (
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/data/validation"
)

var (
	ErrRequiredField    = errors.New("обязательное поле")
	ErrTooShortString   = errors.New("слишком короткая строка")
	ErrTooLongString    = errors.New("слишком длинная строка")
	ErrTooShortPassword = errors.New("слишком короткий пароль")
	ErrTooLongPassword  = errors.New("слишком длинный пароль")
)

// RequiredField non empty string validator
func RequiredField(input string) error {
	if input == "" {
		return ErrRequiredField
	}
	return nil
}

// NewLenRange creates string validator for length.
//
// Ignore limit if is non positive
func NewLenRange(from, to int) func(string) error {
	return func(s string) error {
		if from > 0 && len(s) < from {
			return ErrTooShortString
		}
		if to > 0 && len(s) > to {
			return ErrTooLongString
		}
		return nil
	}
}

// NewPasswordLenRange creates string validator for password length.
//
// Ignore limit if is non positive
func NewPasswordLenRange(from, to int) func(string) error {
	strValidator := NewLenRange(from, to)
	return func(s string) error {
		err := strValidator(s)
		if errors.Is(err, ErrTooShortString) {
			return fmt.Errorf("%w: не менее %d символов", ErrTooShortPassword, from)
		}
		if errors.Is(err, ErrTooLongString) {
			return fmt.Errorf("%w: не более %d символов", ErrTooLongPassword, to)
		}
		return nil
	}
}

func NewPasswordCharsValidator() func(string) error {
	lower := "qwertyuiopasdfghjklzxcvbnm"
	upper := "QWERTYUIOPASDFGHJKLZXCVBNM"
	digits := "1234567890"
	specs := "_-+=!@#$%^&*?"

	all := strings.Split(lower+upper+digits+specs, "")
	allowed := make(map[string]bool, len(all))
	for _, s := range all {
		allowed[s] = true
	}

	return func(input string) error {
		syms := strings.Split(input, "")
		for _, s := range syms {
			if !allowed[s] {
				return fmt.Errorf("не поддерживаемый символ %s", s)
			}
		}

		if !strings.ContainsAny(input, lower) {
			return fmt.Errorf("пароль должен содержать латинские символы в нижнем регистре")
		}
		if !strings.ContainsAny(input, upper) {
			return fmt.Errorf("пароль должен содержать латинские символы в верхнем регистре")
		}
		if !strings.ContainsAny(input, specs) {
			return fmt.Errorf("пароль должен содержать хотя бы один спецсимвол: %s", specs)
		}

		return nil
	}
}

func NewPasswordDefaultValidator() func(string) error {
	return validation.NewAllStrings(
		RequiredField,
		NewPasswordCharsValidator(),
		NewPasswordLenRange(6, 24),
	)
}
