package employee

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
)

type Service interface {
	GetEmployeeInformation(ctx context.Context, req model.GetEmployeeInformationRequest) (error, []model.GetEmployeeInformationResponse)
	GetEmployeeEditInformation(ctx context.Context, req model.GetEmployeeByIdRequest) (error, []model.GetEmployeeByIdResponse)
	GetEmployeeMasterAddress(ctx context.Context, req model.GetEmployeeMasterAddressRequest) (error, []model.GetEmployeeMasterAddressResponse)
	UpdateEmployeeMasterAddress(ctx context.Context, req model.UpdateEmployeeMasterAddressRequest) (error, string)
	CreateEmployeeMasterAddress(ctx context.Context, req model.CreateEmployeeMasterAddressRequest) (error, string)
}
