// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeClient interface {
	GetEmployeeInformation(ctx context.Context, in *GetEmployeeInformationRequest, opts ...grpc.CallOption) (*GetEmployeeInformationListResponse, error)
}

type employeeClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeClient(cc grpc.ClientConnInterface) EmployeeClient {
	return &employeeClient{cc}
}

func (c *employeeClient) GetEmployeeInformation(ctx context.Context, in *GetEmployeeInformationRequest, opts ...grpc.CallOption) (*GetEmployeeInformationListResponse, error) {
	out := new(GetEmployeeInformationListResponse)
	err := c.cc.Invoke(ctx, "/model.Employee/GetEmployeeInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServer is the server API for Employee service.
// All implementations must embed UnimplementedEmployeeServer
// for forward compatibility
type EmployeeServer interface {
	GetEmployeeInformation(context.Context, *GetEmployeeInformationRequest) (*GetEmployeeInformationListResponse, error)
	mustEmbedUnimplementedEmployeeServer()
}

// UnimplementedEmployeeServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServer struct {
}

func (UnimplementedEmployeeServer) GetEmployeeInformation(context.Context, *GetEmployeeInformationRequest) (*GetEmployeeInformationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeInformation not implemented")
}
func (UnimplementedEmployeeServer) mustEmbedUnimplementedEmployeeServer() {}

// UnsafeEmployeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServer will
// result in compilation errors.
type UnsafeEmployeeServer interface {
	mustEmbedUnimplementedEmployeeServer()
}

func RegisterEmployeeServer(s grpc.ServiceRegistrar, srv EmployeeServer) {
	s.RegisterService(&Employee_ServiceDesc, srv)
}

func _Employee_GetEmployeeInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).GetEmployeeInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Employee/GetEmployeeInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).GetEmployeeInformation(ctx, req.(*GetEmployeeInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Employee_ServiceDesc is the grpc.ServiceDesc for Employee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployeeInformation",
			Handler:    _Employee_GetEmployeeInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
