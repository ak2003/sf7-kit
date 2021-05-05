package example

import (
	"context"
	"database/sql"
	"errors"
	"sf7-kit/pkg/example/model/protoc/model"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db     *sql.DB
}

type Repository interface {
	HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (model.HealthCheckResponse, error)
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db:     db,
	}
}

func (r repo) HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (model.HealthCheckResponse, error) {
	// @ select table
	//err := r.db.QueryRow("SELECT email FROM mt_user WHERE id=$1", wording).Scan(&wording)
	//if err != nil {
	//	level.Error(logCreate).Log("err", err)
	//	return "", RepoErr
	//}

	// @ insert
	//var query = `
	//	INSERT INTO mt_user (id, email, password)
	//	VALUES ($1, $2, $3)`
	//
	//_, err := repo.db.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	//if err != nil {
	//	level.Error(logCreate).Log("err", err)
	//	return err
	//}

	var res model.HealthCheckResponse
	res.Wording = "Are You Say, " + req.Wording

	return res, nil
}