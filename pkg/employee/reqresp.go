package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
)

func decodeGetEmployeeInformationReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetEmployeeInformationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmployeeListingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetEmployeeListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmployeeEditInformationReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetEmployeeByIdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmployeeMasterAddressReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetEmployeeMasterAddressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmployeeUpdateAddressReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.UpdateEmployeeMasterAddressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmployeeCreateAddressReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.CreateEmployeeMasterAddressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEmploymentStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetEmploymentStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetJobGradeReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetJobGradeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCityReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetCityRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetAddressTypeReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetAddressTypeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetOwnerStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetOwnerStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetStayStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetStayStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
