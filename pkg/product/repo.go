package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/product/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
)

//var logCreate = logger.MakeLogEntry("product", "RepoProduct")

type Product struct {
	ID          string        `json:"id,omitempty"`
	ProductName string        `json:"product_name,omitempty"`
	CategoryID  int           `json:"category_id,omitempty"`
	BrandID     int           `json:"brand_id,omitempty"`
	Description []Description `json:"description,omitempty"`
	Price       int           `json:"price,omitempty"`
	DiscPrice   int           `json:"disc_price,omitempty"`
	DiscPercent int           `json:"disc_percent,omitempty"`
	Options     []Options     `json:"options,omitempty"`
	Gallery     []string      `json:"gallery,omitempty"`
	SupplierID  int           `json:"supplier_id,omitempty"`
}

type Description struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Options struct {
	Title       string        `json:"title,omitempty"`
	Type        string        `json:"type,omitempty"`
	itemOptions []ItemOptions `json:"item_options,omitempty"`
	IsMandatory bool          `json:"is_mandatory"`
}

type ItemOptions struct {
	Value string `json:"value,omitempty"`
	Price string `json:"price,omitempty"`
}

type Repository interface {
	CreateProduct(ctx context.Context, product Product) (*sql.Tx, error)
	DeleteProduct(ct context.Context, uid string) error
	DetailProduct(ctx context.Context, id string) (*model.ProductDetail, error)
}

type repo struct {
	db     *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db:     db,
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
		//level.Error(logCreate).Log("err", err)
		return nil, err
	}

	// Variant
	v, err = json.Marshal(p.Options)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return nil, err
	}

	// Gallery
	g, err = json.Marshal(p.Gallery)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return nil, err
	}

	var query = `
		INSERT INTO mt_product (id, name, category_id, supplier_id, description, options, gallery, price, disc_price, disc_percent)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	tx, errTx := repo.db.Begin()
	if errTx != nil {
		//level.Error(logCreate).Log("err", errTx)
		return nil, err
	}
	_, err = tx.ExecContext(ctx, query, p.ID, p.ProductName, p.CategoryID, p.SupplierID, string(d), string(v), string(g), p.Price, p.DiscPrice, p.DiscPercent)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return nil, err
	}
	return tx, nil
}

func (repo *repo) DetailProduct(ctx context.Context, id string) (*model.ProductDetail, error) {
	var (
		p       model.ProductDetail
		gallery string
		options string
	)

	err := repo.db.QueryRow("SELECT id,name,gallery,options, price::money::numeric::int8 FROM mt_product WHERE id=$1", id).Scan(&p.Id,&p.ProductName,&gallery, &options, &p.Price)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}

	err = json.Unmarshal([]byte(gallery), &p.Gallery)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}

	err = json.Unmarshal([]byte(options), &p.Options)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}
	return &p, nil
}

func (repo *repo) DeleteProduct(ctx context.Context, id string) error {
	var query = `
	DELETE FROM mt_product where id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		//level.Error(logCreate).Log("err", err)
		return err
	}
	return nil
}
