package product

import (
	"context"
	"fmt"
	"gt-kit/product/model/protoc/model"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) CreateProduct(ctx context.Context, product interface{}) (output interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateProduct", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CreateProduct(ctx, product)
	return
}

func (mw InstrumentingMiddleware) List(ctx context.Context, param *model.ProductId) (output *model.ProductList, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateProduct", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.List(ctx, param)
	return
}