package custom_field

import (
	"context"
	"fmt"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) CheckAddField(ctx context.Context, req *model.AddFieldCheckRequest) (output *model.AddFieldCheckResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CheckAddField", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CheckAddField(ctx, req)
	return
}
