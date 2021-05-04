package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sf7-kit/shared/response"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	LoginUser  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		LoginUser:  makeLoginUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		resp, err := s.CreateUser(ctx, req.Email, req.Password)
		responseBody := response.Body{Message: resp}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}

func makeLoginUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateLoginRequest)
		result, err := s.LoginUser(ctx, req.Email, req.Password)
		data := CreateLoginResponse{
			Token: result,
		}
		responseBody := response.Body{Data: data}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id, req.Token)
		responseBody := response.Body{Data: email}
		return response.CreateResponse{RespBody: responseBody, Err: err}, nil
	}
}
