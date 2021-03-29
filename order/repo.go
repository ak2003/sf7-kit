package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"gt-kit/order/model"
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

func (repo *repo) SaveShoppingCart(ctx context.Context, sc model.ShoppingCart) error {
	i, err := json.Marshal(sc.Items)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return err
	}

	var query = `INSERT INTO tr_shopping_cart (id, user_id, items, total) VALUES ($1, $2, $3, $4)`
	_, err = repo.db.ExecContext(ctx, query, sc.ID, sc.UserID, i, sc.Total)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
}

func (repo *repo) GetShoppingCart(ctx context.Context, id string) (*model.ShoppingCart, error) {
	var (
		sc model.ShoppingCart
		item string
	)

	err := repo.db.QueryRow("SELECT id,user_id,items,total::money::numeric::int8 FROM tr_shopping_cart WHERE id=$1", id).Scan(&sc.ID,&sc.UserID,&item,&sc.Total)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(item), &sc.Items)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return nil, err
	}

	return &sc, nil
}

func (repo *repo) UpdateItemShoppingCart(ctx context.Context, cartId string, itemCart []model.ItemCart, total int64) error {
	ic, err := json.Marshal(itemCart)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return err
	}

	var query = `
	UPDATE tr_shopping_cart SET items=$2, total=$3  where id = $1`
	_, err = repo.db.ExecContext(ctx, query, cartId, ic, total)
	if err != nil {
		level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
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
