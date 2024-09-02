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
	defer os.RemoveAll(dbPath)

	cryptoService := service.NewCryptoService()
	tokenService := service.NewTokenService(tokenRepo, cryptoService)

	tokenDTO := dto.NewTokenDTO(
		"Example-Title",
		"", // TODO: description field unused in Entitie and real database
		"examplePassword",
		"exampleToken",
	)

	if err := tokenService.AddToken(ctx, tokenDTO); err != nil {
		t.Fatal(err)
	}
	decTokenDTO, err := tokenService.GetTokenByTitleDecrypted(ctx, tokenDTO.Title, tokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if decTokenDTO != tokenDTO {
		t.Fatalf("got %+v, want %+v", decTokenDTO, tokenDTO)
	}

	if err = tokenService.AddToken(ctx, tokenDTO); errors.Is(err, entity.ErrTokenExist) {
		t.Fatalf("false correct add existed token: %v", err)
	}

	tokenDTO.Password = "newPassword"
	if err := tokenService.UpdateToken(ctx, tokenDTO); err != nil {
		t.Fatal(err)
	}
	decTokenDTO, err = tokenService.GetTokenByTitleDecrypted(ctx, tokenDTO.Title, tokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if decTokenDTO != tokenDTO {
		t.Fatalf("got %+v, want %+v", decTokenDTO, tokenDTO)
	}

	secondTokenDTO := dto.NewTokenDTO(
		"second token title",
		"",
		"",
		"OpenedToken",
	)
	if err := tokenService.AddToken(ctx, secondTokenDTO); err != nil {
		t.Fatal(err)
	}
	if decDTO, err := tokenService.GetTokenByTitle(ctx, secondTokenDTO.Title); err != nil {
		t.Fatal(err)
	} else if decDTO != secondTokenDTO {
		t.Fatalf("got %+v, want %+v", decDTO, secondTokenDTO)
	}
	decDTO, err := tokenService.GetTokenByTitleDecrypted(ctx, secondTokenDTO.Title, secondTokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if decDTO != secondTokenDTO {
		t.Fatalf("got %+v, want %+v", decDTO, secondTokenDTO)
	}

	dtos, err := tokenService.ListTokens(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(dtos) != 2 {
		t.Fatalf("incorrect tokens list size: got %d, want %d\nlist=%+v", len(dtos), 2, dtos)
	}

	titles, err := tokenService.ListTokensTitles(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, title := range titles {
		wantTitles := []string{tokenDTO.Title, secondTokenDTO.Title}
		if !slices.Contains(wantTitles, title) {
			log.Fatalf("titles list not contains %s", title)
		}
	}

	secondTokenDTO.Password = "now_SettedPassword"
	if err := tokenService.UpdateToken(ctx, secondTokenDTO); err != nil {
		log.Fatal(err)
	}
	decDTO, err = tokenService.GetTokenByTitleDecrypted(ctx, secondTokenDTO.Title, secondTokenDTO.Password)
	if err != nil {
		t.Fatal(err)
	}
	if decDTO != secondTokenDTO {
		t.Fatalf("got %+v, want %+v", decDTO, secondTokenDTO)
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
