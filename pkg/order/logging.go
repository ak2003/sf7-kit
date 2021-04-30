package order

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	model2 "gt-kit/pkg/order/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

func (mw LoggingMiddleware) AddToCart(ctx context.Context, req model2.AddToCartRequest) (output interface{}, err error) {
	var (
		i []byte
	)

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

func (mw LoggingMiddleware) DeleteItemCart(ctx context.Context, req model2.DeleteItemCartRequest) (output *[]model2.ItemCart, err error) {
	var (
		i []byte
	)

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

	output, err = mw.Next.DeleteItemCart(ctx, req)
	return
}

