package employee

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
)

type Service interface {
	GetEmployeeInformation(ctx context.Context, req model.GetEmployeeInformationRequest) (error, []model.GetEmployeeInformationResponse)
	GetEmployeeListing(ctx context.Context, req model.GetEmployeeListingRequest) (error, model.GetEmployeeListingResponse)
	GetEmployeeEditInformation(ctx context.Context, req model.GetEmployeeByIdRequest) (error, []model.GetEmployeeByIdResponse)
	GetEmployeeMasterAddress(ctx context.Context, req model.GetEmployeeMasterAddressRequest) (error, []model.GetEmployeeMasterAddressResponse)
	UpdateEmployeeMasterAddress(ctx context.Context, req model.UpdateEmployeeMasterAddressRequest) (error, string)
	CreateEmployeeMasterAddress(ctx context.Context, req model.CreateEmployeeMasterAddressRequest) (error, string)
	GetCity(ctx context.Context, req model.GetCityRequest) (error, []model.GetCityResponse)
	GetAddressType(ctx context.Context, req model.GetAddressTypeRequest) (error, []model.GetAddressTypeResponse)
	GetOwnerStatus(ctx context.Context, req model.GetOwnerStatusRequest) (error, []model.GetOwnerStatusResponse)
	GetStayStatus(ctx context.Context, req model.GetStayStatusRequest) (error, []model.GetStayStatusResponse)
}
