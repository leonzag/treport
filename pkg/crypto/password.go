package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hashed string, pwd string) bool {
	if pwd == "" && hashed == pwd {
		return true
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	return err == nil
}
