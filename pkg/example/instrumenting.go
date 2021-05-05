package example

import (
	"context"
	"fmt"
	"gitlab.com/dataon1/sf7-kit/pkg/example/model/protoc/model"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (output *model.HealthCheckResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "HealthCheck", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.HealthCheck(ctx, req)
	return
}
