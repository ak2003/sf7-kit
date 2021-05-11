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
	var req model.UpdateEmployeeMasterAddressRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
