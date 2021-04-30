package example

import (
	"context"
	"gt-kit/pkg/example/model/protoc/model"
	"gt-kit/shared/constant"
	"gt-kit/shared/utils/logger"
)

type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s service) HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (interface{}, error) {
	if req.Wording == "" {
		logger.Error(nil, constant.ErrInvalidArgument)
		return nil, constant.ErrInvalidArgument
	}
	res, err := s.repository.HealthCheck(ctx, req)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}
	return res, nil
}
