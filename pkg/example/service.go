package example

import (
	"context"
	"gt-kit/pkg/example/model/protoc/model"
)

type Service interface {
	HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (interface{}, error)
}
