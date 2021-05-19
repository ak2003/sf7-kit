package leave

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
)

type Service interface {
	GetLeaveRequestListing(ctx context.Context, req model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse)
	GetLeaveRequestFilterListing(ctx context.Context, req model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse)
	GetDataTypeOfLeave(ctx context.Context, req model.GetDataTypeOfLeaveReq) (error, []model.GetDataTypeOfLeaveResponse)
	GetDataRequestFor(ctx context.Context, req model.GetDataRequestForReq) (error, []model.GetDataRequestForResponse)
	GetDataRemainingLeave(ctx context.Context, req model.GetDataRemainingLeaveReq) (error, []model.GetDataRemainingLeaveResponse)
	CreateLeaveRequest(ctx context.Context, req model.CreateLeaveRequestReq) (error, string)
	CreateLeaveRequestForm(ctx context.Context, req model.CreateLeaveRequestFormReq) (error, string)
}
