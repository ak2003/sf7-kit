package product

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/product/model/protoc/model"
)

type Service interface {
	CreateProduct(ctx context.Context, product interface{}) (interface{}, error)
	DetailProduct(ctx context.Context, param *model.ProductId) (*model.ProductDetail, error)
}