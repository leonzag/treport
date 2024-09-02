package crypto_test

import (
	"testing"

	"github.com/leonzag/treport/pkg/crypto"
)

func TestHashPassword(t *testing.T) {
	original := "examplePassword"
	hashed, err := crypto.HashPassword(original)
	if err != nil {
		t.Fatal(err)
	}

	if hashed == original {
		t.Fatal("password was not hashed")
	}
}

func TestCheckPassword(t *testing.T) {
	original := "examplePassword"
	hashed, err := crypto.HashPassword(original)
	if err != nil {
		t.Fatal(err)
	}
	if !crypto.CheckPassword(hashed, original) {
		t.Fatal("incorrect password check")
	}
	if crypto.CheckPassword(hashed, "incorrectPassword") {
		t.Fatal("incorrect password check")
	}
}
