package product

import (
	"context"
	"database/sql"
	"gt-kit/product/model/protoc/model"
)

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
