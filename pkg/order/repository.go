package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"sf7-kit/pkg/order/model"
	"sf7-kit/shared/utils/logger"

)

type Repository interface {
	SaveShoppingCart(ctx context.Context, sc model.ShoppingCart) error
	GetShoppingCart(ctx context.Context, id string) (*model.ShoppingCart, error)
	UpdateItemShoppingCart(ctx context.Context, id string, itemCart []model.ItemCart, total int64) error
	DeleteProduct(ct context.Context, uid string) error
}

type repo struct {
	db     *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db:     db,
	}
}

func (repo *repo) SaveShoppingCart(ctx context.Context, sc model.ShoppingCart) error {
	i, err := json.Marshal(sc.Items)
	if err != nil {
		logger.Error(nil, err)
		return err
	}

	var query = `INSERT INTO tr_shopping_cart (id, user_id, items, total) VALUES ($1, $2, $3, $4)`
	_, err = repo.db.ExecContext(ctx, query, sc.ID, sc.UserID, i, sc.Total)
	if err != nil {
		logger.Error(nil, err)
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
		logger.Error(nil, err)
		return nil, err
	}

	err = json.Unmarshal([]byte(item), &sc.Items)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}

	return &sc, nil
}

func (repo *repo) UpdateItemShoppingCart(ctx context.Context, cartId string, itemCart []model.ItemCart, total int64) error {
	ic, err := json.Marshal(itemCart)
	if err != nil {
		logger.Error(nil, err)
		return err
	}

	var query = `
	UPDATE tr_shopping_cart SET items=$2, total=$3  where id = $1`
	_, err = repo.db.ExecContext(ctx, query, cartId, ic, total)
	if err != nil {
		logger.Error(nil, err)
		return err
	}
	return nil
}

func (repo *repo) DeleteProduct(ctx context.Context, id string) error {
	var query = `
	DELETE FROM mt_product where id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error(nil, err)
		return err
	}
	return nil
}
