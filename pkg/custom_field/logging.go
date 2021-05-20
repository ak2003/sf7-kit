package custom_field

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
	"time"
)

type LoggingMiddleware struct {
	Next Service
}

func (mw LoggingMiddleware) CheckAddField(ctx context.Context, req *model.AddFieldCheckRequest) (output *model.AddFieldCheckResponse, err error) {
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

	return mw.Next.CheckAddField(ctx, req)
}
