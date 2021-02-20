package order

import (
	"context"
	"time"
)

type ShoppingCart struct {
	ID        string
	UserID    string
	Items     []ItemCart
	MetaData  interface{}
	Total     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ItemCart struct {
	ProductID       string
	Image           string
	Qty             int
	Price           int64
	OptionsItemCart []OptionsItemCart
}

type OptionsItemCart struct {
	Title        string
	ItemSelected string
	Price        int64
}

type Repository interface {
	SaveShoppingCart(ctx context.Context, sc ShoppingCart) error
	GetShoppingCart(ctx context.Context, id string) (*ShoppingCart, error)
	UpdateItemShoppingCart(ctx context.Context, id string, itemCart []ItemCart, total int64) error
	DeleteProduct(ct context.Context, uid string) error
}
