// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/grpc_employee/model/employee.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEmployeeInformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId  string `protobuf:"bytes,1,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
	EmployeeId string `protobuf:"bytes,2,opt,name=EmployeeId,proto3" json:"EmployeeId,omitempty"`
	Language   string `protobuf:"bytes,3,opt,name=Language,proto3" json:"Language,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Page       string `protobuf:"bytes,5,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit      string `protobuf:"bytes,6,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetEmployeeInformationRequest) Reset() {
	*x = GetEmployeeInformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeInformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeInformationRequest) ProtoMessage() {}

func (x *GetEmployeeInformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeInformationRequest.ProtoReflect.Descriptor instead.
func (*GetEmployeeInformationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_employee_model_employee_proto_rawDescGZIP(), []int{0}
}

func (x *GetEmployeeInformationRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *GetEmployeeInformationRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *GetEmployeeInformationRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetEmployeeInformationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetEmployeeInformationRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetEmployeeInformationRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type GetEmployeeInformationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeName          string  `protobuf:"bytes,1,opt,name=EmployeeName,proto3" json:"EmployeeName,omitempty"`
	EmployeeId            string  `protobuf:"bytes,2,opt,name=EmployeeId,proto3" json:"EmployeeId,omitempty"`
	EmployeeNo            string  `protobuf:"bytes,3,opt,name=EmployeeNo,proto3" json:"EmployeeNo,omitempty"`
	EmployeePos           string  `protobuf:"bytes,4,opt,name=EmployeePos,proto3" json:"EmployeePos,omitempty"`
	EmployeePhoneExt      *string `protobuf:"bytes,5,opt,name=EmployeePhoneExt,proto3,oneof" json:"EmployeePhoneExt,omitempty"`
	EmployeeDept          string  `protobuf:"bytes,6,opt,name=EmployeeDept,proto3" json:"EmployeeDept,omitempty"`
	EmployeeStartDate     string  `protobuf:"bytes,7,opt,name=EmployeeStartDate,proto3" json:"EmployeeStartDate,omitempty"`
	EmployeeGrade         string  `protobuf:"bytes,8,opt,name=EmployeeGrade,proto3" json:"EmployeeGrade,omitempty"`
	EmployeeStatus        string  `protobuf:"bytes,9,opt,name=EmployeeStatus,proto3" json:"EmployeeStatus,omitempty"`
	EmployeeEmail         string  `protobuf:"bytes,10,opt,name=EmployeeEmail,proto3" json:"EmployeeEmail,omitempty"`
	EmployeePhoto         string  `protobuf:"bytes,11,opt,name=EmployeePhoto,proto3" json:"EmployeePhoto,omitempty"`
	EmployeePhone         *string `protobuf:"bytes,12,opt,name=EmployeePhone,proto3,oneof" json:"EmployeePhone,omitempty"`
	EmployeeMaritalStatus string  `protobuf:"bytes,13,opt,name=EmployeeMaritalStatus,proto3" json:"EmployeeMaritalStatus,omitempty"`
	EmployeeEndDate       string  `protobuf:"bytes,14,opt,name=EmployeeEndDate,proto3" json:"EmployeeEndDate,omitempty"`
	EmployeeGenderName    string  `protobuf:"bytes,15,opt,name=EmployeeGenderName,proto3" json:"EmployeeGenderName,omitempty"`
	EmployeeReqFlag       string  `protobuf:"bytes,16,opt,name=EmployeeReqFlag,proto3" json:"EmployeeReqFlag,omitempty"`
	EmployeeGenderCode    string  `protobuf:"bytes,17,opt,name=EmployeeGenderCode,proto3" json:"EmployeeGenderCode,omitempty"`
}

func (x *GetEmployeeInformationResponse) Reset() {
	*x = GetEmployeeInformationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeInformationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeInformationResponse) ProtoMessage() {}

func (x *GetEmployeeInformationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeInformationResponse.ProtoReflect.Descriptor instead.
func (*GetEmployeeInformationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_employee_model_employee_proto_rawDescGZIP(), []int{1}
}

func (x *GetEmployeeInformationResponse) GetEmployeeName() string {
	if x != nil {
		return x.EmployeeName
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeNo() string {
	if x != nil {
		return x.EmployeeNo
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeePos() string {
	if x != nil {
		return x.EmployeePos
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeePhoneExt() string {
	if x != nil && x.EmployeePhoneExt != nil {
		return *x.EmployeePhoneExt
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeDept() string {
	if x != nil {
		return x.EmployeeDept
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeStartDate() string {
	if x != nil {
		return x.EmployeeStartDate
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeGrade() string {
	if x != nil {
		return x.EmployeeGrade
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeStatus() string {
	if x != nil {
		return x.EmployeeStatus
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeEmail() string {
	if x != nil {
		return x.EmployeeEmail
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeePhoto() string {
	if x != nil {
		return x.EmployeePhoto
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeePhone() string {
	if x != nil && x.EmployeePhone != nil {
		return *x.EmployeePhone
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeMaritalStatus() string {
	if x != nil {
		return x.EmployeeMaritalStatus
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeEndDate() string {
	if x != nil {
		return x.EmployeeEndDate
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeGenderName() string {
	if x != nil {
		return x.EmployeeGenderName
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeReqFlag() string {
	if x != nil {
		return x.EmployeeReqFlag
	}
	return ""
}

func (x *GetEmployeeInformationResponse) GetEmployeeGenderCode() string {
	if x != nil {
		return x.EmployeeGenderCode
	}
	return ""
}

type GetEmployeeInformationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GetEmployeeInformationResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetEmployeeInformationListResponse) Reset() {
	*x = GetEmployeeInformationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeInformationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeInformationListResponse) ProtoMessage() {}

func (x *GetEmployeeInformationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_employee_model_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeInformationListResponse.ProtoReflect.Descriptor instead.
func (*GetEmployeeInformationListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_employee_model_employee_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmployeeInformationListResponse) GetList() []*GetEmployeeInformationResponse {
	if x != nil {
		return x.List
	}
	return nil
}

var File_pkg_grpc_employee_model_employee_proto protoreflect.FileDescriptor

var file_pkg_grpc_employee_model_employee_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0xbb, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xff, 0x05,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x50, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x2f, 0x0a, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x45, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x45, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x44, 0x65, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x65, 0x70, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x29, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x15, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x0a, 0x12, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x2e, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x45, 0x78, 0x74, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x5f, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0x77, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x6b, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_grpc_employee_model_employee_proto_rawDescOnce sync.Once
	file_pkg_grpc_employee_model_employee_proto_rawDescData = file_pkg_grpc_employee_model_employee_proto_rawDesc
)

func file_pkg_grpc_employee_model_employee_proto_rawDescGZIP() []byte {
	file_pkg_grpc_employee_model_employee_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_employee_model_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_employee_model_employee_proto_rawDescData)
	})
	return file_pkg_grpc_employee_model_employee_proto_rawDescData
}

var file_pkg_grpc_employee_model_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_grpc_employee_model_employee_proto_goTypes = []interface{}{
	(*GetEmployeeInformationRequest)(nil),      // 0: model.GetEmployeeInformationRequest
	(*GetEmployeeInformationResponse)(nil),     // 1: model.GetEmployeeInformationResponse
	(*GetEmployeeInformationListResponse)(nil), // 2: model.GetEmployeeInformationListResponse
}
var file_pkg_grpc_employee_model_employee_proto_depIdxs = []int32{
	1, // 0: model.GetEmployeeInformationListResponse.list:type_name -> model.GetEmployeeInformationResponse
	0, // 1: model.Employee.GetEmployeeInformation:input_type -> model.GetEmployeeInformationRequest
	2, // 2: model.Employee.GetEmployeeInformation:output_type -> model.GetEmployeeInformationListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpc_employee_model_employee_proto_init() }
func file_pkg_grpc_employee_model_employee_proto_init() {
	if File_pkg_grpc_employee_model_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_employee_model_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeInformationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_employee_model_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeInformationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_employee_model_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeInformationListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_grpc_employee_model_employee_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_employee_model_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_employee_model_employee_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_employee_model_employee_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_employee_model_employee_proto_msgTypes,
	}.Build()
	File_pkg_grpc_employee_model_employee_proto = out.File
	file_pkg_grpc_employee_model_employee_proto_rawDesc = nil
	file_pkg_grpc_employee_model_employee_proto_goTypes = nil
	file_pkg_grpc_employee_model_employee_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
type EmployeeServer interface {
	GetEmployeeInformation(context.Context, *GetEmployeeInformationRequest) (*GetEmployeeInformationListResponse, error)
}

// UnimplementedEmployeeServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServer struct {
}

func (*UnimplementedEmployeeServer) GetEmployeeInformation(context.Context, *GetEmployeeInformationRequest) (*GetEmployeeInformationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeInformation not implemented")
}

func RegisterEmployeeServer(s *grpc.Server, srv EmployeeServer) {
	s.RegisterService(&_Employee_serviceDesc, srv)
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

var _Employee_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployeeInformation",
			Handler:    _Employee_GetEmployeeInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc_employee/model/employee.proto",
}