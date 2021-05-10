package product

import (
	"context"
	"fmt"
	"time"

	"gitlab.dataon.com/gophers/sf7-kit/internal/product/model/protoc/model"

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

func (mw InstrumentingMiddleware) DetailProduct(ctx context.Context, param *model.ProductId) (output *model.ProductDetail, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateProduct", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.DetailProduct(ctx, param)
	return
}
