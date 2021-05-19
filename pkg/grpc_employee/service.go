package grpc_employee

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee/model/protoc/model"
)

type Service interface {
	GetEmployeeInformation(ctx context.Context, req *model.GetEmployeeInformationRequest) (*model.GetEmployeeInformationListResponse, error)
}
