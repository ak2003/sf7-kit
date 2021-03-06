// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/custom_field/model/AdditionalField.proto

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

type AddFieldCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId    string `protobuf:"bytes,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
}

func (x *AddFieldCheckRequest) Reset() {
	*x = AddFieldCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldCheckRequest) ProtoMessage() {}

func (x *AddFieldCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldCheckRequest.ProtoReflect.Descriptor instead.
func (*AddFieldCheckRequest) Descriptor() ([]byte, []int) {
	return file_pkg_custom_field_model_AdditionalField_proto_rawDescGZIP(), []int{0}
}

func (x *AddFieldCheckRequest) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *AddFieldCheckRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type AddFieldCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Label       string       `protobuf:"bytes,2,opt,name=Label,proto3" json:"Label,omitempty"`
	Type        string       `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Name        string       `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	DefaultName string       `protobuf:"bytes,5,opt,name=DefaultName,proto3" json:"DefaultName,omitempty"`
	IsMandatory bool         `protobuf:"varint,6,opt,name=IsMandatory,proto3" json:"IsMandatory,omitempty"`
	Data        []*DataField `protobuf:"bytes,7,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *AddFieldCheck) Reset() {
	*x = AddFieldCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldCheck) ProtoMessage() {}

func (x *AddFieldCheck) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldCheck.ProtoReflect.Descriptor instead.
func (*AddFieldCheck) Descriptor() ([]byte, []int) {
	return file_pkg_custom_field_model_AdditionalField_proto_rawDescGZIP(), []int{1}
}

func (x *AddFieldCheck) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddFieldCheck) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AddFieldCheck) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddFieldCheck) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddFieldCheck) GetDefaultName() string {
	if x != nil {
		return x.DefaultName
	}
	return ""
}

func (x *AddFieldCheck) GetIsMandatory() bool {
	if x != nil {
		return x.IsMandatory
	}
	return false
}

func (x *AddFieldCheck) GetData() []*DataField {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddFieldCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddFieldCheck []*AddFieldCheck `protobuf:"bytes,1,rep,name=addFieldCheck,proto3" json:"addFieldCheck,omitempty"`
}

func (x *AddFieldCheckResponse) Reset() {
	*x = AddFieldCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFieldCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFieldCheckResponse) ProtoMessage() {}

func (x *AddFieldCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFieldCheckResponse.ProtoReflect.Descriptor instead.
func (*AddFieldCheckResponse) Descriptor() ([]byte, []int) {
	return file_pkg_custom_field_model_AdditionalField_proto_rawDescGZIP(), []int{2}
}

func (x *AddFieldCheckResponse) GetAddFieldCheck() []*AddFieldCheck {
	if x != nil {
		return x.AddFieldCheck
	}
	return nil
}

type DataField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *DataField) Reset() {
	*x = DataField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataField) ProtoMessage() {}

func (x *DataField) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_custom_field_model_AdditionalField_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataField.ProtoReflect.Descriptor instead.
func (*DataField) Descriptor() ([]byte, []int) {
	return file_pkg_custom_field_model_AdditionalField_proto_rawDescGZIP(), []int{3}
}

func (x *DataField) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pkg_custom_field_model_AdditionalField_proto protoreflect.FileDescriptor

var file_pkg_custom_field_model_AdditionalField_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x4c, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x4d, 0x61, 0x6e, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73, 0x4d, 0x61,
	0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a,
	0x15, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x22, 0x31, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x5f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x64, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_custom_field_model_AdditionalField_proto_rawDescOnce sync.Once
	file_pkg_custom_field_model_AdditionalField_proto_rawDescData = file_pkg_custom_field_model_AdditionalField_proto_rawDesc
)

func file_pkg_custom_field_model_AdditionalField_proto_rawDescGZIP() []byte {
	file_pkg_custom_field_model_AdditionalField_proto_rawDescOnce.Do(func() {
		file_pkg_custom_field_model_AdditionalField_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_custom_field_model_AdditionalField_proto_rawDescData)
	})
	return file_pkg_custom_field_model_AdditionalField_proto_rawDescData
}

var file_pkg_custom_field_model_AdditionalField_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_custom_field_model_AdditionalField_proto_goTypes = []interface{}{
	(*AddFieldCheckRequest)(nil),  // 0: model.AddFieldCheckRequest
	(*AddFieldCheck)(nil),         // 1: model.AddFieldCheck
	(*AddFieldCheckResponse)(nil), // 2: model.AddFieldCheckResponse
	(*DataField)(nil),             // 3: model.DataField
}
var file_pkg_custom_field_model_AdditionalField_proto_depIdxs = []int32{
	3, // 0: model.AddFieldCheck.Data:type_name -> model.DataField
	1, // 1: model.AddFieldCheckResponse.addFieldCheck:type_name -> model.AddFieldCheck
	0, // 2: model.AdditionalField.CheckAddField:input_type -> model.AddFieldCheckRequest
	2, // 3: model.AdditionalField.CheckAddField:output_type -> model.AddFieldCheckResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_custom_field_model_AdditionalField_proto_init() }
func file_pkg_custom_field_model_AdditionalField_proto_init() {
	if File_pkg_custom_field_model_AdditionalField_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_custom_field_model_AdditionalField_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFieldCheckRequest); i {
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
		file_pkg_custom_field_model_AdditionalField_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFieldCheck); i {
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
		file_pkg_custom_field_model_AdditionalField_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFieldCheckResponse); i {
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
		file_pkg_custom_field_model_AdditionalField_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataField); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_custom_field_model_AdditionalField_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_custom_field_model_AdditionalField_proto_goTypes,
		DependencyIndexes: file_pkg_custom_field_model_AdditionalField_proto_depIdxs,
		MessageInfos:      file_pkg_custom_field_model_AdditionalField_proto_msgTypes,
	}.Build()
	File_pkg_custom_field_model_AdditionalField_proto = out.File
	file_pkg_custom_field_model_AdditionalField_proto_rawDesc = nil
	file_pkg_custom_field_model_AdditionalField_proto_goTypes = nil
	file_pkg_custom_field_model_AdditionalField_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdditionalFieldClient is the client API for AdditionalField service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionalFieldClient interface {
	CheckAddField(ctx context.Context, in *AddFieldCheckRequest, opts ...grpc.CallOption) (*AddFieldCheckResponse, error)
}

type additionalFieldClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionalFieldClient(cc grpc.ClientConnInterface) AdditionalFieldClient {
	return &additionalFieldClient{cc}
}

func (c *additionalFieldClient) CheckAddField(ctx context.Context, in *AddFieldCheckRequest, opts ...grpc.CallOption) (*AddFieldCheckResponse, error) {
	out := new(AddFieldCheckResponse)
	err := c.cc.Invoke(ctx, "/model.AdditionalField/CheckAddField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionalFieldServer is the server API for AdditionalField service.
type AdditionalFieldServer interface {
	CheckAddField(context.Context, *AddFieldCheckRequest) (*AddFieldCheckResponse, error)
}

// UnimplementedAdditionalFieldServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionalFieldServer struct {
}

func (*UnimplementedAdditionalFieldServer) CheckAddField(context.Context, *AddFieldCheckRequest) (*AddFieldCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAddField not implemented")
}

func RegisterAdditionalFieldServer(s *grpc.Server, srv AdditionalFieldServer) {
	s.RegisterService(&_AdditionalField_serviceDesc, srv)
}

func _AdditionalField_CheckAddField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFieldCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionalFieldServer).CheckAddField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AdditionalField/CheckAddField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionalFieldServer).CheckAddField(ctx, req.(*AddFieldCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdditionalField_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.AdditionalField",
	HandlerType: (*AdditionalFieldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAddField",
			Handler:    _AdditionalField_CheckAddField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/custom_field/model/AdditionalField.proto",
}
