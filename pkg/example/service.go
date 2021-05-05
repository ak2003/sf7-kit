package example

import (
	"context"
	"gitlab.com/dataon1/sf7-kit/pkg/example/model/protoc/model"
)

type Service interface {
	HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (*model.HealthCheckResponse, error)
}
