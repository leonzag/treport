package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/leonzag/treport/internal/domain/entity"
)

func (r *tokenRepo) Add(ctx context.Context, token *entity.Token) error {
	r.Lock()
	defer r.Unlock()

	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `
	INSERT INTO tokens (
		title, password, token
	) VALUES (?, ?, ?);
	`
	_, err = tx.ExecContext(ctx, stmt, token.Title, token.Password, token.Token)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *tokenRepo) Get(ctx context.Context, title string) (*entity.Token, error) {
	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	token := &entity.Token{}

	stmt := `
	SELECT title, password, token
	FROM tokens
	WHERE title=?;
	`
	row := tx.QueryRowContext(ctx, stmt, title)
	err = row.Scan(&token.Title, &token.Password, &token.Token)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrTokenNotFound
	}
	if err != nil {
		return nil, err
	}
	return token, tx.Commit()
}

func (r *tokenRepo) List(ctx context.Context) ([]*entity.Token, error) {
	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	q := `
	SELECT
		title, password, token
	FROM
		tokens
	`
	rows, err := tx.QueryContext(ctx, q)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*entity.Token{}, entity.ErrTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	tokens := []*entity.Token{}

	for rows.Next() {
		token := &entity.Token{}
		err := rows.Scan(&token.Title, &token.Password, &token.Token)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	if rows.Err() != nil {
		return tokens, rows.Err()
	}
	return tokens, tx.Commit()
}

func (r *tokenRepo) Update(ctx context.Context, token *entity.Token) error {
	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
	UPDATE
		tokens
	SET
		token=?,
		password=?
	WHERE
		title=?
	`
	_, err = tx.ExecContext(ctx, q, token.Token, token.Password, token.Title)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *tokenRepo) Delete(ctx context.Context, token *entity.Token) error {
	_, err := r.Get(ctx, token.Title)
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
	DELETE FROM
		tokens
	WHERE
		title=?
	`
	if _, err = tx.ExecContext(ctx, q, token.Title); err != nil {
		return err
	}
	return tx.Commit()
}
