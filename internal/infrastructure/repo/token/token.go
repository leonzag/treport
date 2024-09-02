package token

import (
	"github.com/leonzag/treport/internal/infrastructure/repo"
	"github.com/leonzag/treport/internal/infrastructure/repo/token/sqlite"
	"github.com/leonzag/treport/pkg/database"
)

func NewTokenSQliteRepo(db database.DB) repo.TokenRepo {
	return sqlite.NewTokenRepo(db)
}
