package product

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	model2 "gt-kit/pkg/product/model/protoc/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

func (mw LoggingMiddleware) CreateProduct(ctx context.Context, product interface{}) (output interface{}, err error) {
	var i []byte

	i, err = json.Marshal(product)
	if err != nil {
		return nil, err
	}

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateProduct",
			"input", i,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.CreateProduct(ctx, product)
	return
}

func (mw LoggingMiddleware) DetailProduct(ctx context.Context, param *model2.ProductId) (output *model2.ProductDetail, err error) {
	var i []byte

	i, err = json.Marshal(param)
	if err != nil {
		return nil, err
	}

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "ListProduct",
			"input", i,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.DetailProduct(ctx, param)
	return
}
