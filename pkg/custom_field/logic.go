package custom_field

import (
	"context"
	"encoding/json"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
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

func (s service) CheckAddField(ctx context.Context, req *model.AddFieldCheckRequest) (interface{}, error) {
	p := &model.AddFieldCheckResponse{}
	emptyData := make([]string, 0)
	res, err := s.repository.AddFieldCheck(ctx, req)
	if err != nil {
		logger.Error(nil, err)
		return emptyData, nil
	}

	err = json.Unmarshal([]byte(res.AdditionalFields), &p.AddFieldCheck)
	if err != nil {
		logger.Error(nil, err)
		return nil, err
	}

	return p.AddFieldCheck, nil
}
