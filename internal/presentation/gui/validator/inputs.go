package validator

import "fmt"

func RequiredField(input string) error {
	if input == "" {
		return fmt.Errorf("обязательное поле")
	}
	return nil
}
