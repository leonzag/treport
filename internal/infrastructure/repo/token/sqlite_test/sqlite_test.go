package sqlite_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/leonzag/treport/internal/domain/entity"
	tokenRepo "github.com/leonzag/treport/internal/infrastructure/repo/token/sqlite"
	"github.com/leonzag/treport/pkg/database/sqlite"
)

var dbPath = "./token_test.sqlite"

func TestSQLiteTokenRepo(t *testing.T) {
	database, err := sqlite.New(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dbPath)

	ctx := context.TODO()
	repo := tokenRepo.NewTokenRepo(database)

	if !repo.IsInited() {
		if err := repo.Init(ctx); err != nil {
			t.Fatal(fmt.Errorf("failed sqlite db init: %w", err))
		}
	}
	defer os.RemoveAll(dbPath)

	t.Log("test <repo.List> with empty db")
	tokensList, err := repo.List(ctx)
	if err == nil && len(tokensList) != 0 {
		t.Fatalf("unexpectedly got tokensList=%+v", tokensList)
	} else if err != nil && !errors.Is(err, entity.ErrTokenNotFound) {
		msg := fmt.Errorf("want: <ErrTokenNotFound> got: %w", err)
		t.Fatal(msg)
	}

	t.Log("test <repo.Add>")
	wantToken := entity.Token{
		Title:    "TestTokenTitle-1",
		Password: "",
		Token:    "someOpenToken",
	}
	_, err = repo.Add(ctx, &wantToken)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("test <repo.Get>")
	gotTokenPtr, err := repo.Get(ctx, wantToken.Title)
	if err != nil {
		t.Fatal(err)
	}
	if gotTokenPtr == nil || *gotTokenPtr != wantToken {
		t.Fatalf("want: %+v \ngot: %+v", wantToken, gotTokenPtr)
	}

	t.Log("test <repo.Add> (2)")
	_, err = repo.Add(ctx, &entity.Token{
		Title:    "TestTokenTitle-2",
		Password: "examplePassword",
		Token:    "someEncryptedToken",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("test <repo.List>")
	tokensList, err = repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wantLen := 2
	if ln := len(tokensList); ln != wantLen {
		t.Fatalf("incorrect list length (got %d, want %d)\ntokensList:%+v", ln, wantLen, tokensList)
	}

	t.Log("test <repo.Update>")
	wantToken.Token = "someOpenTokenChanged"
	wantToken.Password = "nowWithPassword"

	_, err = repo.Update(ctx, &wantToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test <repo.Get> on updated token")
	gotTokenPtr, err = repo.Get(ctx, wantToken.Title)
	if err != nil {
		t.Fatal(err)
	}
	if gotTokenPtr == nil || *gotTokenPtr != wantToken {
		t.Fatalf("want: %+v \ngot: %+v", wantToken, gotTokenPtr)
	}

	t.Log("test <repo.Delete>")
	err = repo.Delete(ctx, &wantToken)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("test <repo.List> (2)")
	tokensList, err = repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wantLen = 1
	if ln := len(tokensList); ln != wantLen {
		t.Fatalf("incorrect list length (got %d, want %d)\ntokensList:%v", ln, wantLen, tokensList)
	}

	t.Log("test <repo.Get> on non-existent token")
	unexpectedToken, err := repo.Get(ctx, wantToken.Title)
	if !errors.Is(err, entity.ErrTokenNotFound) {
		msg := fmt.Errorf("unexpectedly token=%+v err=%w", unexpectedToken, err)
		t.Fatal(msg)
	}

	t.Log("test <repo.Close>")
	err = repo.Close()
	if err != nil {
		t.Fatal(err)
	}
}
