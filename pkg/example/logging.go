package example

import (
	"context"
	"sf7-kit/pkg/example/model/protoc/model"
	"sf7-kit/shared/utils/logger"
	"time"
)

type LoggingMiddleware struct {
	Next Service
}

func (mw LoggingMiddleware) HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (output *model.HealthCheckResponse, err error) {
	msg := "Incoming Request"
	defer func(begin time.Time) {
		fields := make(map[string]interface{})
		fields["input"] = req
		fields["output"] = output
		fields["err"] = err
		if err != nil {
			logger.Error(fields, msg)
			return
		}
		logger.Info(fields, msg)
	}(time.Now())

	return mw.Next.HealthCheck(ctx, req)
}
