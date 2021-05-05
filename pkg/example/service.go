package example

import (
	"context"
	"sf7-kit/pkg/example/model/protoc/model"
)

type Service interface {
	HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (*model.HealthCheckResponse, error)
}
