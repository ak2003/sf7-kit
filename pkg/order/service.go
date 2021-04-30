package order

import (
	"context"
	"gt-kit/pkg/order/model"
)

type Service interface {
	AddToCart(ctx context.Context, addToCart model.AddToCartRequest) (interface{}, error)
	DeleteItemCart(ctx context.Context, params model.DeleteItemCartRequest) (*[]model.ItemCart, error)
}