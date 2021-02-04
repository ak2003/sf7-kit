package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, product interface{}) (interface{}, error)
	//CreateUser(ctx context.Context, email string, password string) (string, error)
	//GetUser(ctx context.Context, id string, tokenString string) (string, error)
	//LoginUser(ctx context.Context, username string, password string) (string, error)
}