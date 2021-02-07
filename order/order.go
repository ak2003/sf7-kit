package order

import (
	"context"
	"database/sql"
	"time"
)

type ShoppingCart struct {
	ID        string
	UserID    string
	Items     ItemCart
	MetaData  interface{}
	Total     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ItemCart struct {
	ProductID   string
	Sku         string
	Image       string
	Qty         int
	Price       float64
	Description string
}

type Repository interface {
	SaveShoppingCart(ctx context.Context, sc ShoppingCart) (*sql.Tx, error)
	DeleteProduct(ct context.Context, uid string) error
}
