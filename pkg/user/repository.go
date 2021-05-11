package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	model "gitlab.dataon.com/gophers/sf7-kit/pkg/user/model/user"
)

var RepoErr = errors.New("unable to handle repo request")

//var logCreate = logger.MakeLogEntry("user","RepoUser")

type repo struct {
	dbSlave  *sqlx.DB
	dbMaster *sqlx.DB
	logger   log.Logger
}

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, email string) (string, string, error)
	CheckEmail(ctx context.Context, username string) (int, error)
}

func NewRepo(dbSlave, dbMaster *sqlx.DB, logger log.Logger) Repository {
	return &repo{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
		logger:   log.With(logger, "repo", "sql"),
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

	_, err := repo.dbMaster.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.dbSlave.QueryRow("SELECT email FROM mt_user WHERE id=$1", id).Scan(&email)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return "", RepoErr
	}

	return email, nil
}

func (repo *repo) LoginUser(ctx context.Context, username string) (string, string, error) {
	var (
		email    string
		password string
	)
	err := repo.dbSlave.QueryRow("SELECT email, password FROM mt_user WHERE email=$1", username).Scan(&email, &password)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return "", "", RepoErr
	}

	return email, password, nil
}

func (repo *repo) CheckEmail(ctx context.Context, username string) (int, error) {
	var emailCount int
	err := repo.dbSlave.QueryRow("SELECT count (email) as emailCount FROM mt_user WHERE email=$1", username).Scan(&emailCount)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return 0, RepoErr
	}

	return emailCount, nil
}
