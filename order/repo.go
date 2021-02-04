package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"gt-kit/shared/utils/logger"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to handle repo request")
var logCreate = logger.MakeLogEntry("product", "RepoProduct")

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

func (repo *repo) CreateProduct(ctx context.Context, p Product) (*sql.Tx, error) {
	var (
		err error
		d []byte
		v []byte
		g []byte
	)

	// Description
	d, err = json.Marshal(p.Description)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}

	// Variant
	v, err = json.Marshal(p.Variant)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}

	// Gallery
	g, err = json.Marshal(p.Gallery)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}

	var query = `
		INSERT INTO mt_product (id, name, category_id, supplier_id, description, variant, gallery, price, disc_price, disc_percent)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	tx, errTx := repo.db.Begin()
	if errTx != nil {
		level.Error(logCreate).Log("err", errTx)
		return nil, err
	}
	_, err = tx.ExecContext(ctx, query, p.ID, p.ProductName, p.CategoryID, p.SupplierID, string(d), string(v), string(g), p.Price, p.DiscPrice, p.DiscPercent)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}
	return tx, nil
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
