package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gt-kit/product/model/protoc/model"
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
	v, err = json.Marshal(p.Options)
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
		INSERT INTO mt_product (id, name, category_id, supplier_id, description, options, gallery, price, disc_price, disc_percent)
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

func (repo *repo) DetailProduct(ctx context.Context, id string) (*model.ProductDetail, error) {
	var (
		p model.ProductDetail
		gallery string
		options string
		a interface{}
	)

	err := repo.db.QueryRow("SELECT id,name,gallery,options FROM mt_product WHERE id=$1", id).Scan(&p.Id,&p.ProductName,&gallery, &options)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}
	json.Unmarshal([]byte(gallery), &p.Gallery)
	json.Unmarshal([]byte(options), &a)
	fmt.Printf("%+v", a)
	//json.Unmarshal([]byte(options), &p.Gallery)
	return &p, nil
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
