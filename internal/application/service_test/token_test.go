package servicetest

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/application/interfaces"
	"github.com/leonzag/treport/internal/application/service"
	"github.com/leonzag/treport/internal/domain/entity"
	tokenRepo "github.com/leonzag/treport/internal/infrastructure/repo/token/sqlite"
	"github.com/leonzag/treport/pkg/database/sqlite"
)

func TestTokenService(t *testing.T) {
	ctx := context.TODO()
	dbPath := filepath.Join(os.TempDir(), "./test_token.sqlite")

	tokenRepo := createTempDb(ctx, dbPath)
	t.Cleanup(func() {
		os.RemoveAll(dbPath)
	})

	cryptoService := service.NewCryptoService()
	tokenService := service.NewTokenService(tokenRepo, cryptoService)

	tokenDTO := dto.TokenRequestDTO{
		Title:       "Example-Title",
		Description: "",
		Password:    "examplePassword",
		Token:       "exampleToken",
	}
	want := &entity.TokenDecrypted{
		Title: tokenDTO.Title,
		Token: tokenDTO.Token,
	}
	wantTitles := []string{tokenDTO.Title}

	if _, err := tokenService.AddToken(ctx, tokenDTO); err != nil {
		t.Fatal(err)
	}
	token, err := tokenService.GetTokenByTitle(ctx, tokenDTO.Title)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := tokenService.DecryptToken(token, tokenDTO.Password)
	if err != nil {
		t.Fatalf("failed decrypt: %s", err.Error())
	}
	if dec.Title != want.Title || dec.Token != want.Token {
		t.Fatalf("got %+v, want %+v", dec, want)
	}

	_, err = tokenService.AddToken(ctx, tokenDTO)
	if !errors.Is(err, entity.ErrTokenExist) {
		t.Fatalf("false correct add existed token: %v", err)
	}

	tokenDTO.Password = "newPassword"
	if _, err := tokenService.UpdateToken(ctx, tokenDTO); err != nil {
		t.Fatal(err)
	}

	token, err = tokenService.GetTokenByTitle(ctx, tokenDTO.Title)
	if err != nil {
		t.Fatal(err)
	}
	dec, err = tokenService.DecryptToken(token, tokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}

	if dec.Title != tokenDTO.Title || dec.Token != tokenDTO.Token {
		t.Fatalf("got %+v, want %+v", dec, want)
	}

	tokenDTO = dto.TokenRequestDTO{
		Title:       "second token title",
		Description: "",
		Password:    "",
		Token:       "OpenedToken",
	}
	want.Title = tokenDTO.Title
	want.Token = tokenDTO.Token
	wantTitles = append(wantTitles, tokenDTO.Title)

	token, err = tokenService.AddToken(ctx, tokenDTO)
	if err != nil {
		t.Fatal(err)
	}
	if token.Title != want.Title || token.Token != want.Token {
		t.Fatalf("got %+v, want %+v", dec, want)
	}

	token, err = tokenService.GetTokenByTitle(ctx, tokenDTO.Title)
	if err != nil {
		t.Fatal(err)
	}

	dec, err = tokenService.DecryptToken(token, tokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if dec.Title != want.Title || dec.Token != want.Token {
		t.Fatalf("got %+v, want %+v", dec, want)
	}

	tokens, err := tokenService.ListTokens(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(tokens) != 2 {
		t.Fatalf("incorrect tokens list size: got %d, want %d\nlist=%+v", len(tokens), 2, tokens)
	}

	titles, err := tokenService.ListTokensTitles(ctx)
	if err != nil {
		t.Fatal(err)
	}
	slices.Sort(titles)
	slices.Sort(wantTitles)
	if !slices.Equal(titles, wantTitles) {
		t.Fatalf("not equal: got %+v, want %+v", titles, wantTitles)
	}

	tokenDTO.Password = "now_SettedPassword"
	want.Token = tokenDTO.Token

	token, err = tokenService.UpdateToken(ctx, tokenDTO)
	if err != nil {
		log.Fatal(err)
	}
	dec, err = tokenService.DecryptToken(token, tokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if dec.Title != want.Title || dec.Token != want.Token {
		t.Fatalf("got %+v, want %+v", dec, want)
	}
}

func createTempDb(ctx context.Context, dbPath string) interfaces.TokenRepo {
	db, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	tokenRepo := tokenRepo.NewTokenRepo(db)
	if !tokenRepo.IsInited() {
		if err := tokenRepo.Init(ctx); err != nil {
			log.Fatal(err)
		}
	}
	return tokenRepo
}
