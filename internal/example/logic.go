package example

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/internal/example/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/constant"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
)

type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s service) HealthCheck(ctx context.Context, req *model.HealthCheckRequest) (*model.HealthCheckResponse, error) {
	if req.Wording == "" {
		logger.Error(nil, constant.ErrInvalidArgument)
		return nil, constant.ErrInvalidArgument
	}
	res, err := s.repository.HealthCheck(ctx, req)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}
	return &res, nil
}
