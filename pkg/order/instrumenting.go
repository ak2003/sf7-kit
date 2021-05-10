package order

import (
	"context"
	"fmt"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/order/model"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}


func (mw InstrumentingMiddleware) AddToCart(ctx context.Context, req model.AddToCartRequest) (output interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddToCart", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.AddToCart(ctx, req)
	return
}

func (mw InstrumentingMiddleware)  DeleteItemCart(ctx context.Context, req model.DeleteItemCartRequest) (output *[]model.ItemCart, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddToCart", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.DeleteItemCart(ctx, req)
	return
}
