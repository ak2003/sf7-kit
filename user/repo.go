package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	var query = `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}

	return email, nil
}

func (repo *repo) LoginUser(ctx context.Context, username string) (string, string, error) {
	var (
		email string
		password string
	)
	err := repo.db.QueryRow("SELECT email, password FROM users WHERE email=$1", username).Scan(&email, &password)
	if err != nil {
		return "", "", RepoErr
	}

	return email, password, nil
}

func (repo *repo) CheckEmail(ctx context.Context, username string) (int, error) {
	var emailCount int
	err := repo.db.QueryRow("SELECT count(email) as emailCount FROM users WHERE email=$1", username).Scan(&emailCount)
	if err != nil {
		return 0, RepoErr
	}

	return emailCount, nil
}
