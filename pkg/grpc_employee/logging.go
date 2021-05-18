package grpc_employee

import (
	"context"
	"time"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
)

type LoggingMiddleware struct {
	Next Service
}

func (mw LoggingMiddleware) GetEmployeeInformation(ctx context.Context, req *model.GetEmployeeInformationRequest) (output []*model.GetEmployeeInformationResponse, err error) {
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

	return mw.Next.GetEmployeeInformation(ctx, req)
}
