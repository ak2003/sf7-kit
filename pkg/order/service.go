package order

import (
	"context"
	"gitlab.com/dataon1/sf7-kit/pkg/order/model"
)

type Service interface {
	AddToCart(ctx context.Context, addToCart model.AddToCartRequest) (interface{}, error)
	DeleteItemCart(ctx context.Context, params model.DeleteItemCartRequest) (*[]model.ItemCart, error)
}