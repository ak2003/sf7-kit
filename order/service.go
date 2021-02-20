package order

import (
	"context"
	"gt-kit/order/model"
)

type Service interface {
	AddToCart(ctx context.Context, addToCart model.AddToCartRequest) (interface{}, error)
	DeleteItemCart(ctx context.Context, params model.DeleteItemCartRequest) (*[]model.ItemCart, error)
	//CreateUser(ctx context.Context, email string, password string) (string, error)
	//GetUser(ctx context.Context, id string, tokenString string) (string, error)
	//LoginUser(ctx context.Context, username string, password string) (string, error)
}