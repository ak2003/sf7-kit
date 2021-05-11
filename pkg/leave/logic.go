package leave

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
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

func (s service) GetLeaveRequestListing(ctx context.Context, param model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse) {
	//logDetail := logger.MakeLogEntry("product", "DetailProduct")
	//level.Info(logDetail).Log("param-id", param.Id)

	err, dp := s.repository.GetLeaveRequestListing(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetLeaveRequestFilterListing(ctx context.Context, param model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse) {
	//logDetail := logger.MakeLogEntry("product", "DetailProduct")
	//level.Info(logDetail).Log("param-id", param.Id)

	err, dp := s.repository.GetLeaveRequestFilterListing(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}
