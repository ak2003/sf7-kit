package employee

import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
)

type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	var srvc = service{
		repository: rep,
	}
	return &srvc
}

func (s service) CreateEmployeeMasterAddress(ctx context.Context, param model.CreateEmployeeMasterAddressRequest) (error, string) {
	err, dp := s.repository.CreateEmployeeMasterAddress(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) UpdateEmployeeMasterAddress(ctx context.Context, param model.UpdateEmployeeMasterAddressRequest) (error, string) {
	err, dp := s.repository.UpdateEmployeeMasterAddress(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetEmployeeInformation(ctx context.Context, param model.GetEmployeeInformationRequest) (error, []model.GetEmployeeInformationResponse) {
	//logDetail := logger.MakeLogEntry("product", "DetailProduct")
	//level.Info(logDetail).Log("param-id", param.Id)

	err, dp := s.repository.GetEmployeeInformation(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetEmployeeEditInformation(ctx context.Context, param model.GetEmployeeByIdRequest) (error, []model.GetEmployeeByIdResponse) {
	err, dp := s.repository.GetEmployeeEditInformation(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetEmployeeMasterAddress(ctx context.Context, param model.GetEmployeeMasterAddressRequest) (error, []model.GetEmployeeMasterAddressResponse) {
	err, dp := s.repository.GetEmployeeMasterAddress(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetCity(ctx context.Context, param model.GetCityRequest) (error, []model.GetCityResponse) {
	err, dp := s.repository.GetCity(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetAddressType(ctx context.Context, param model.GetAddressTypeRequest) (error, []model.GetAddressTypeResponse) {
	err, dp := s.repository.GetAddressType(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetOwnerStatus(ctx context.Context, param model.GetOwnerStatusRequest) (error, []model.GetOwnerStatusResponse) {
	err, dp := s.repository.GetOwnerStatus(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}

func (s service) GetStayStatus(ctx context.Context, param model.GetStayStatusRequest) (error, []model.GetStayStatusResponse) {
	err, dp := s.repository.GetStayStatus(ctx, param)
	if err != nil {
		logger.Error(nil, err)
		return nil, dp
	}

	return nil, dp
}
