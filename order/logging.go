package order

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

//func (mw LoggingMiddleware) CreateProduct(ctx context.Context, product interface{}) (output interface{}, err error) {
//	var i []byte
//
//	i, err = json.Marshal(product)
//	if err != nil {
//		return nil, err
//	}
//
//	defer func(begin time.Time) {
//		_ = mw.Logger.Log(
//			"method", "CreateProduct",
//			"input", i,
//			"output", output,
//			"err", err,
//			"took", time.Since(begin),
//		)
//	}(time.Now())
//
//	output, err = mw.Next.CreateProduct(ctx, product)
//	return
//}

func (mw LoggingMiddleware) AddToCart(ctx context.Context, req AddToCartRequest) (output interface{}, err error) {
	var i []byte

	i, err = json.Marshal(req)
	if err != nil {
		return nil, err
	}

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "AddToCart",
			"input", i,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.AddToCart(ctx, req)
	return
}

