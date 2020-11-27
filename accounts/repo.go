package accounts

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

var RepoErr = errors.New("Unabele to handle Repo Request")

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r *repo) CreateUser(ctx context.Context, user User) error {
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	sql := `
	INSERT INTO users (id, email, password)
	VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	sql := "SELECT email FROM users WHERE id=$1"

	if err := r.db.QueryRow(sql, id).Scan(&email); err != nil {
		return "", RepoErr
	}

	return email, nil
}
