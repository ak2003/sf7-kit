package order

import (
	"context"
	"gt-kit/order/model"
)

type Repository interface {
	SaveShoppingCart(ctx context.Context, sc model.ShoppingCart) error
	GetShoppingCart(ctx context.Context, id string) (*model.ShoppingCart, error)
	UpdateItemShoppingCart(ctx context.Context, id string, itemCart []model.ItemCart, total int64) error
	DeleteProduct(ct context.Context, uid string) error
}
