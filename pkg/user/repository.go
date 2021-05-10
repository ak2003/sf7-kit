package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
	model "gitlab.dataon.com/gophers/sf7-kit/pkg/user/model/user"
)

var RepoErr = errors.New("unable to handle repo request")
//var logCreate = logger.MakeLogEntry("user","RepoUser")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, email string) (string, string, error)
	CheckEmail(ctx context.Context, username string) (int, error)
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) error {
	var query = `
		INSERT INTO mt_user (id, email, password)
		VALUES ($1, $2, $3)`
	if user.Email == "" || user.Password == "" {
		//level.Info(logCreate).Log("msg", "username, password is empty")
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM mt_user WHERE id=$1", id).Scan(&email)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return "", RepoErr
	}

	return email, nil
}

func (repo *repo) LoginUser(ctx context.Context, username string) (string, string, error) {
	var (
		email string
		password string
	)
	err := repo.db.QueryRow("SELECT email, password FROM mt_user WHERE email=$1", username).Scan(&email, &password)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return "", "", RepoErr
	}

	return email, password, nil
}

func (repo *repo) CheckEmail(ctx context.Context, username string) (int, error) {
	var emailCount int
	err := repo.db.QueryRow("SELECT count (email) as emailCount FROM mt_user WHERE email=$1", username).Scan(&emailCount)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return 0, RepoErr
	}

	return emailCount, nil
}
