package order

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/internal/order/model"
)

type Service interface {
	AddToCart(ctx context.Context, addToCart model.AddToCartRequest) (interface{}, error)
	DeleteItemCart(ctx context.Context, params model.DeleteItemCartRequest) (*[]model.ItemCart, error)
}