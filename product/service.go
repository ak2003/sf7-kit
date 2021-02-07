package product

import (
	"context"
	"gt-kit/product/model/protoc/model"
)

type Service interface {
	CreateProduct(ctx context.Context, product interface{}) (interface{}, error)
	List(ctx context.Context, param *model.ProductId) (*model.ProductList, error)
	//CreateUser(ctx context.Context, email string, password string) (string, error)
	//GetUser(ctx context.Context, id string, tokenString string) (string, error)
	//LoginUser(ctx context.Context, username string, password string) (string, error)
}