package leave

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
)

type Service interface {
	GetLeaveRequestListing(ctx context.Context, req model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse)
	GetLeaveRequestFilterListing(ctx context.Context, req model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse)
	GetDataRequestFor(ctx context.Context, req model.GetDataRequestForReq) (error, []model.GetDataRequestForResponse)
}
