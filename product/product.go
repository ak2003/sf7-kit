package product

import (
	"context"
	"database/sql"
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
	Variant     []Variant     `json:"variant,omitempty"`
	Gallery     []string      `json:"gallery,omitempty"`
	SupplierID  int           `json:"supplier_id,omitempty"`
}

type Description struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Variant struct {
	Title   string       `json:"title,omitempty"`
	Type    string       `json:"type,omitempty"`
	Options []VariantOpt `json:"options,omitempty"`
}

type VariantOpt struct {
	Title string `json:"title,omitempty"`
	Price string `json:"price,omitempty"`
}

type Repository interface {
	CreateProduct(ctx context.Context, product Product) (*sql.Tx, error)
	DeleteProduct(ct context.Context, uid string) error
}
