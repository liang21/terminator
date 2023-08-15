// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: api/test/v1/module.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TestModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectId int64  `protobuf:"varint,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *TestModule) Reset() {
	*x = TestModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestModule) ProtoMessage() {}

func (x *TestModule) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestModule.ProtoReflect.Descriptor instead.
func (*TestModule) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{0}
}

func (x *TestModule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestModule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestModule) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type CreateTestModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // the name of string must be between 5 and 50 character
	ProjectId int64  `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateTestModuleRequest) Reset() {
	*x = CreateTestModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestModuleRequest) ProtoMessage() {}

func (x *CreateTestModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestModuleRequest.ProtoReflect.Descriptor instead.
func (*CreateTestModuleRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTestModuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTestModuleRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type CreateTestModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestModule *TestModule `protobuf:"bytes,1,opt,name=TestModule,proto3" json:"TestModule,omitempty"`
}

func (x *CreateTestModuleReply) Reset() {
	*x = CreateTestModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestModuleReply) ProtoMessage() {}

func (x *CreateTestModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestModuleReply.ProtoReflect.Descriptor instead.
func (*CreateTestModuleReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTestModuleReply) GetTestModule() *TestModule {
	if x != nil {
		return x.TestModule
	}
	return nil
}

type UpdateTestModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // the name of string must be between 5 and 50 character
	ProjectId int64  `protobuf:"varint,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *UpdateTestModuleRequest) Reset() {
	*x = UpdateTestModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestModuleRequest) ProtoMessage() {}

func (x *UpdateTestModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestModuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateTestModuleRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTestModuleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTestModuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTestModuleRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type UpdateTestModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestModule *TestModule `protobuf:"bytes,1,opt,name=TestModule,proto3" json:"TestModule,omitempty"`
}

func (x *UpdateTestModuleReply) Reset() {
	*x = UpdateTestModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestModuleReply) ProtoMessage() {}

func (x *UpdateTestModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestModuleReply.ProtoReflect.Descriptor instead.
func (*UpdateTestModuleReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTestModuleReply) GetTestModule() *TestModule {
	if x != nil {
		return x.TestModule
	}
	return nil
}

type DeleteTestModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTestModuleRequest) Reset() {
	*x = DeleteTestModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestModuleRequest) ProtoMessage() {}

func (x *DeleteTestModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestModuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestModuleRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTestModuleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTestModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTestModuleReply) Reset() {
	*x = DeleteTestModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestModuleReply) ProtoMessage() {}

func (x *DeleteTestModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestModuleReply.ProtoReflect.Descriptor instead.
func (*DeleteTestModuleReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{6}
}

type GetTestModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTestModuleRequest) Reset() {
	*x = GetTestModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestModuleRequest) ProtoMessage() {}

func (x *GetTestModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestModuleRequest.ProtoReflect.Descriptor instead.
func (*GetTestModuleRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{7}
}

func (x *GetTestModuleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTestModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestModule *TestModule `protobuf:"bytes,1,opt,name=TestModule,proto3" json:"TestModule,omitempty"`
}

func (x *GetTestModuleReply) Reset() {
	*x = GetTestModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestModuleReply) ProtoMessage() {}

func (x *GetTestModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestModuleReply.ProtoReflect.Descriptor instead.
func (*GetTestModuleReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{8}
}

func (x *GetTestModuleReply) GetTestModule() *TestModule {
	if x != nil {
		return x.TestModule
	}
	return nil
}

type ListTestModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTestModuleRequest) Reset() {
	*x = ListTestModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestModuleRequest) ProtoMessage() {}

func (x *ListTestModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestModuleRequest.ProtoReflect.Descriptor instead.
func (*ListTestModuleRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{9}
}

func (x *ListTestModuleRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTestModuleRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTestModuleRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTestModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int64         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Results []*TestModule `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListTestModuleReply) Reset() {
	*x = ListTestModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_module_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestModuleReply) ProtoMessage() {}

func (x *ListTestModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_module_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestModuleReply.ProtoReflect.Descriptor instead.
func (*ListTestModuleReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_module_proto_rawDescGZIP(), []int{10}
}

func (x *ListTestModuleReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListTestModuleReply) GetResults() []*TestModule {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_test_v1_module_proto protoreflect.FileDescriptor

var file_api_test_v1_module_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x79, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0x92, 0x41,
	0x0b, 0x32, 0x09, 0xe9, 0xa1, 0xb5, 0xe5, 0xa4, 0xa7, 0xe5, 0xb0, 0x8f, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32,
	0x0c, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe9, 0xa1, 0xb5, 0xe7, 0xa0, 0x81, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x69, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x32, 0xdb, 0x05, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f,
	0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f,
	0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x1a, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8f,
	0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x86, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x42, 0x66, 0x92, 0x41, 0x53, 0x12, 0x2a, 0x0a, 0x20, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x41, 0x50, 0x49, 0x32, 0x06, 0x76, 0x31, 0x2e, 0x30,
	0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_test_v1_module_proto_rawDescOnce sync.Once
	file_api_test_v1_module_proto_rawDescData = file_api_test_v1_module_proto_rawDesc
)

func file_api_test_v1_module_proto_rawDescGZIP() []byte {
	file_api_test_v1_module_proto_rawDescOnce.Do(func() {
		file_api_test_v1_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_test_v1_module_proto_rawDescData)
	})
	return file_api_test_v1_module_proto_rawDescData
}

var file_api_test_v1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_test_v1_module_proto_goTypes = []interface{}{
	(*TestModule)(nil),              // 0: terminator.api.test.v1.TestModule
	(*CreateTestModuleRequest)(nil), // 1: terminator.api.test.v1.CreateTestModuleRequest
	(*CreateTestModuleReply)(nil),   // 2: terminator.api.test.v1.CreateTestModuleReply
	(*UpdateTestModuleRequest)(nil), // 3: terminator.api.test.v1.UpdateTestModuleRequest
	(*UpdateTestModuleReply)(nil),   // 4: terminator.api.test.v1.UpdateTestModuleReply
	(*DeleteTestModuleRequest)(nil), // 5: terminator.api.test.v1.DeleteTestModuleRequest
	(*DeleteTestModuleReply)(nil),   // 6: terminator.api.test.v1.DeleteTestModuleReply
	(*GetTestModuleRequest)(nil),    // 7: terminator.api.test.v1.GetTestModuleRequest
	(*GetTestModuleReply)(nil),      // 8: terminator.api.test.v1.GetTestModuleReply
	(*ListTestModuleRequest)(nil),   // 9: terminator.api.test.v1.ListTestModuleRequest
	(*ListTestModuleReply)(nil),     // 10: terminator.api.test.v1.ListTestModuleReply
}
var file_api_test_v1_module_proto_depIdxs = []int32{
	0,  // 0: terminator.api.test.v1.CreateTestModuleReply.TestModule:type_name -> terminator.api.test.v1.TestModule
	0,  // 1: terminator.api.test.v1.UpdateTestModuleReply.TestModule:type_name -> terminator.api.test.v1.TestModule
	0,  // 2: terminator.api.test.v1.GetTestModuleReply.TestModule:type_name -> terminator.api.test.v1.TestModule
	0,  // 3: terminator.api.test.v1.ListTestModuleReply.results:type_name -> terminator.api.test.v1.TestModule
	1,  // 4: terminator.api.test.v1.TestModuleService.CreateTestModule:input_type -> terminator.api.test.v1.CreateTestModuleRequest
	3,  // 5: terminator.api.test.v1.TestModuleService.UpdateTestModule:input_type -> terminator.api.test.v1.UpdateTestModuleRequest
	5,  // 6: terminator.api.test.v1.TestModuleService.DeleteTestModule:input_type -> terminator.api.test.v1.DeleteTestModuleRequest
	7,  // 7: terminator.api.test.v1.TestModuleService.GetTestModule:input_type -> terminator.api.test.v1.GetTestModuleRequest
	9,  // 8: terminator.api.test.v1.TestModuleService.ListTestModule:input_type -> terminator.api.test.v1.ListTestModuleRequest
	2,  // 9: terminator.api.test.v1.TestModuleService.CreateTestModule:output_type -> terminator.api.test.v1.CreateTestModuleReply
	4,  // 10: terminator.api.test.v1.TestModuleService.UpdateTestModule:output_type -> terminator.api.test.v1.UpdateTestModuleReply
	6,  // 11: terminator.api.test.v1.TestModuleService.DeleteTestModule:output_type -> terminator.api.test.v1.DeleteTestModuleReply
	8,  // 12: terminator.api.test.v1.TestModuleService.GetTestModule:output_type -> terminator.api.test.v1.GetTestModuleReply
	10, // 13: terminator.api.test.v1.TestModuleService.ListTestModule:output_type -> terminator.api.test.v1.ListTestModuleReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_test_v1_module_proto_init() }
func file_api_test_v1_module_proto_init() {
	if File_api_test_v1_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_test_v1_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestModule); i {
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
		file_api_test_v1_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestModuleRequest); i {
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
		file_api_test_v1_module_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestModuleReply); i {
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
		file_api_test_v1_module_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestModuleRequest); i {
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
		file_api_test_v1_module_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestModuleReply); i {
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
		file_api_test_v1_module_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestModuleRequest); i {
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
		file_api_test_v1_module_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestModuleReply); i {
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
		file_api_test_v1_module_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestModuleRequest); i {
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
		file_api_test_v1_module_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestModuleReply); i {
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
		file_api_test_v1_module_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestModuleRequest); i {
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
		file_api_test_v1_module_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestModuleReply); i {
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
			RawDescriptor: file_api_test_v1_module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_test_v1_module_proto_goTypes,
		DependencyIndexes: file_api_test_v1_module_proto_depIdxs,
		MessageInfos:      file_api_test_v1_module_proto_msgTypes,
	}.Build()
	File_api_test_v1_module_proto = out.File
	file_api_test_v1_module_proto_rawDesc = nil
	file_api_test_v1_module_proto_goTypes = nil
	file_api_test_v1_module_proto_depIdxs = nil
}
