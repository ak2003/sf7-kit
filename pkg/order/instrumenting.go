package order

import (
	"context"
	"fmt"
	model2 "gt-kit/pkg/order/model"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}


func (mw InstrumentingMiddleware) AddToCart(ctx context.Context, req model2.AddToCartRequest) (output interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddToCart", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.AddToCart(ctx, req)
	return
}

func (mw InstrumentingMiddleware)  DeleteItemCart(ctx context.Context, req model2.DeleteItemCartRequest) (output *[]model2.ItemCart, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddToCart", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.DeleteItemCart(ctx, req)
	return
}
