package grpc_employee

import (
	"context"
	"fmt"
	"time"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee/model/protoc/model"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) GetEmployeeInformation(ctx context.Context, req *model.GetEmployeeInformationRequest) (output []*model.GetEmployeeInformationResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "HealthCheck", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetEmployeeInformation(ctx, req)
	return
}
