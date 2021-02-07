package order

import (
	"context"
	"database/sql"
	"errors"
	"gt-kit/shared/utils/logger"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to handle repo request")
var logCreate = logger.MakeLogEntry("order", "RepoOrder")

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

func (repo *repo) SaveShoppingCart(ctx context.Context, sc ShoppingCart) (*sql.Tx, error) {

	return nil, nil
}

func (repo *repo) DeleteProduct(ctx context.Context, id string) error {
	var query = `
	DELETE FROM mt_product where id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
}
