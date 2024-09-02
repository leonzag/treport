package sqlite

import (
	"context"
	"database/sql"
	"sync"

	define "github.com/leonzag/treport/internal/infrastructure/repo"
	"github.com/leonzag/treport/pkg/database"
)

// check interface impl.
var _ define.TokenRepo = new(tokenRepo)

var txOpts = &sql.TxOptions{}

type tokenRepo struct {
	db database.DB
	mu *sync.RWMutex
}

func NewTokenRepo(db database.DB) *tokenRepo {
	return &tokenRepo{
		db: db,
		mu: &sync.RWMutex{},
	}
}

func (r *tokenRepo) Init(ctx context.Context) error {
	r.Lock()
	defer r.Unlock()

	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tableStmt := `
	CREATE TABLE IF NOT EXISTS tokens (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
		title    TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		token    TEXT NOT NULL
	);`
	if _, err = tx.ExecContext(ctx, tableStmt); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *tokenRepo) IsInited() bool {
	// return r.db.IsInited()
	return false
}

func (r *tokenRepo) Close() error {
	return r.db.Close()
}

func (r *tokenRepo) Lock() {
	r.mu.Lock()
}

func (r *tokenRepo) Unlock() {
	r.mu.Unlock()
}

func (r *tokenRepo) RLock() {
	r.mu.RLock()
}

func (r *tokenRepo) RUnlock() {
	r.mu.RUnlock()
}
